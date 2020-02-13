package unix

import (
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

//
// The following PollServer implementation offers non-blocking I/O capabilities
// to applications that interact with generic file-descriptors. For network
// specific I/O there are known implementations providing similar functionality.
//
// The following API is offered as part of this PollServer implementation:
//
// StartWaitRead(): Place caller goroutine in 'standby' mode till incoming
// traffic is detected over the respective file-descriptor (POLLIN revent
// received).
//
// StartWaitWrite(): Place caller goroutine in 'standby' mode till outgoing
// traffic can be sent through the respective file-descriptor (POLLOUT revent
// received).
//
// StopWait(): Interrupts pollserver event-loop to update the list of file
// descriptors to poll on. This 'interruption' process relies on the use of a
// unidirectional pipe whose receiving-end is being added to the list of file
// descriptors to monitor. This voids the need to define an explicit 'timeout'
// interval during polling attempts.
//

type pollActionType uint8

const (
	CREATE_POLL_REQUEST pollActionType = iota
	DELETE_POLL_REQUEST
)

// Defines the poll action to execute.
type pollAction struct {
	fd          int32          // fd to poll
	actionType  pollActionType // action type
	waitReadCh  chan error     // channel to block on during read operations
	waitWriteCh chan error     // channel to block on during write operations
}

func newPollAction(fd int32, action pollActionType) *pollAction {
	return &pollAction{
		fd:          fd,
		actionType:  action,
		waitReadCh:  make(chan error),
		waitWriteCh: make(chan error),
	}
}

// General pollserver struct.
type PollServer struct {
	sync.RWMutex
	pollActionMap map[int32]*pollAction // 'fd' to 'pollAction' map
	pollActionCh  chan *pollAction      // buffered channel through which pollActions arrive
	wakeupReader  *os.File              // wake-up pipe -- writing end
	wakeupWriter  *os.File              // wake-up pipe -- receiving end
}

func NewPollServer() (*PollServer, error) {
	var err error

	// PollServer struct initialization. Notice that pollActionCh must be buffered
	// to allow incoming pollActions to be injected by pollserver clients *before*
	// the wakeup signal interrupts the polling cycle (see comments below as part
	// of pushPollAction method).
	ps := &PollServer{
		pollActionMap: make(map[int32]*pollAction),
		pollActionCh:  make(chan *pollAction, 1),
	}

	// Initialize pollserver's wakeup pipe.
	ps.wakeupReader, ps.wakeupWriter, err = os.Pipe()
	if err != nil {
		logrus.Errorf("Unable to initialize wakeup pipe in pollserver (%v)", err)
		return nil, err
	}

	// Add wakeup pipe's received-end to the map of fds to monitor.
	ps.pollActionMap[int32(ps.wakeupReader.Fd())] = nil

	go ps.run()

	return ps, nil
}

func (ps *PollServer) StartWaitRead(fd int32) error {

	ps.RLock()
	pa, ok := ps.pollActionMap[fd]
	ps.RUnlock()

	if !ok {
		pa = newPollAction(fd, CREATE_POLL_REQUEST)
	}

	if err := ps.pushPollAction(pa); err != nil {
		return err
	}

	// Block till a POLLIN event is received for this fd.
	err := <-pa.waitReadCh

	return err
}

func (ps *PollServer) StartWaitWrite(fd int32) error {

	ps.RLock()
	pa, ok := ps.pollActionMap[fd]
	ps.RUnlock()

	if !ok {
		pa = newPollAction(fd, CREATE_POLL_REQUEST)
	}

	if err := ps.pushPollAction(pa); err != nil {
		return err
	}

	// Block till a POLLOUT event is received for this fd.
	err := <-pa.waitWriteCh

	return err
}

func (ps *PollServer) StopWait(fd int32) error {

	ps.RLock()
	pa, ok := ps.pollActionMap[fd]
	ps.RUnlock()

	if !ok {
		logrus.Errorf("No poll-request to eliminate for fd %d", fd)
		return fmt.Errorf("No poll-request to eliminate for fd %d", fd)
	}
	pa.actionType = DELETE_POLL_REQUEST

	if err := ps.pushPollAction(pa); err != nil {
		return err
	}

	return nil
}

// Runs pollserver event-loop.
func (ps *PollServer) run() {

	for {
		// Build file-descriptor slice out of the map that tracks all the polling
		// actions to process. No need to acquire rlock here as this is the very
		// goroutine (and the only one) that modifies pollServer internal structs.
		i := 0
		pfds := make([]unix.PollFd, len(ps.pollActionMap))
		for k := range ps.pollActionMap {
			pfds[i].Fd = k
			pfds[i].Events = unix.POLLIN
			i++
		}

		// Initiating polling attempt. Notice that no timeout interval is passed.
		n, err := unix.Poll(pfds, -1)
		if err != nil && err != syscall.EINTR {
			break
		}
		if n <= 0 {
			continue
		}

		// Iterate through all fds to evaluate i/o activity.
		for _, pfd := range pfds {

			if pfd.Revents&unix.POLLHUP == unix.POLLHUP {
				logrus.Debugf("pollserver: POLLHUP event received on fd %d", pfd.Fd)
				continue
			}

			if pfd.Revents&unix.POLLNVAL == unix.POLLNVAL {
				logrus.Debugf("pollserver: POLLNVAL event received on fd %d", pfd.Fd)
				continue
			}

			if pfd.Revents&unix.POLLIN == unix.POLLIN {
				if pfd.Fd == int32(ps.wakeupReader.Fd()) {
					logrus.Debugf("pollserver: WAKEUP event received on fd %d", pfd.Fd)
					ps.processWakeupEvent()
					continue
				}

				logrus.Debugf("pollserver: POLLIN event received on fd %d", pfd.Fd)
				if err := ps.processPollInEvent(&pfd); err != nil {
					logrus.Debugf("pollserver: error POLLIN event received on fd %d", pfd.Fd)
					break
				}
			}

			if pfd.Revents&unix.POLLOUT == unix.POLLOUT {
				logrus.Debugf("pollserver: POLLOUT event received on fd %d", pfd.Fd)
				if err := ps.processPollOutEvent(&pfd); err != nil {
					logrus.Debugf("pollserver: error POLLOUT event received on fd %d", pfd.Fd)
					break
				}
			}
		}
	}
}

// Processes out-of-band 'wakeup' events generated by pollserver clients.
func (ps *PollServer) processWakeupEvent() {

	// TODO: Define this literal in a global var and document its rationale.
	var buf [100]byte
	ps.wakeupReader.Read(buf[0:])

	// Collect received pollAction and add it to the pollActionMap
	pa := <-ps.pollActionCh

	switch pa.actionType {

	case CREATE_POLL_REQUEST:
		ps.Lock()
		ps.pollActionMap[pa.fd] = pa
		ps.Unlock()

	case DELETE_POLL_REQUEST:
		ps.Lock()
		oldPa, ok := ps.pollActionMap[pa.fd]
		if !ok {
			ps.Unlock()
			return
		}
		delete(ps.pollActionMap, pa.fd)
		ps.Unlock()

		// Send an error back to the pollserver client to wake him up from his
		// nap.
		closedFdError := fmt.Errorf("Interrupted poll operation: closed file-descriptor")
		oldPa.waitReadCh <- closedFdError
	}
}

// Processes 'pollin' events received through one of the monitored fds.
func (ps *PollServer) processPollInEvent(pfd *unix.PollFd) error {

	pa, ok := ps.pollActionMap[pfd.Fd]
	if !ok {
		return fmt.Errorf("pollserver: POLLIN event received on unknown fd %v",
			pfd.Fd)
	}

	// Delete fd from pollAction map.
	ps.Lock()
	delete(ps.pollActionMap, pfd.Fd)
	ps.Unlock()

	// Wake-up pollserver client.
	pa.waitReadCh <- nil

	return nil
}

// Processes 'pollout' events received through one of the monitored fds.
func (ps *PollServer) processPollOutEvent(pfd *unix.PollFd) error {

	pa, ok := ps.pollActionMap[pfd.Fd]
	if !ok {
		return fmt.Errorf("pollserver: POLLOUT event received on unknown fd %v",
			pfd.Fd)
	}

	// Delete fd from action map.
	ps.Lock()
	delete(ps.pollActionMap, pfd.Fd)
	ps.Unlock()

	// Wakeup pollserver client.
	pa.waitReadCh <- nil

	return nil
}

// Method pushes incoming pollActions into the pollserver's event-loop. Notice
// that the order of instructions in this method is critical: we must first
// send the pollAction, and then proceed to generate the wakeup signal. If we
// were to invert the order, no incoming state (pollAction) could have been
// received by the time that the pollserver awakes, which would end up with
// the pollserver going back to sleep right before the pollAction arrives.
func (ps *PollServer) pushPollAction(pa *pollAction) error {

	ps.pollActionCh <- pa

	// Write into pollserver's wakeupWriter end to interrupt poll() event loop.
	var buf [1]byte
	ps.wakeupWriter.Write(buf[0:])

	return nil
}
