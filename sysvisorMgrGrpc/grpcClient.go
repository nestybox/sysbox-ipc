// Client-side gRPC interface for the sysvisor manager daemon

package sysvisorMgrGrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/nestybox/sysvisor-ipc/sysvisorMgrGrpc/protobuf"
	"google.golang.org/grpc"
)

func unixConnect(addr string, t time.Duration) (net.Conn, error) {
	unixAddr, err := net.ResolveUnixAddr("unix", grpcSockAddr)
	conn, err := net.DialUnix("unix", nil, unixAddr)
	return conn, err
}

// connect establishes grpc connection to the sysvisorMgr daemon.
func connect() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grpcSockAddr, grpc.WithInsecure(), grpc.WithDialer(unixConnect))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Registers a container with sysvisor-mgr
func Register(id string) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RegisterReq{
		Id: id,
	}

	_, err = ch.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke Register via grpc: %v", err)
	}

	return nil
}

// Unregisters a container with sysvisor-mgr
func Unregister(id string) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UnregisterReq{
		Id: id,
	}

	_, err = ch.Unregister(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke Unregister via grpc: %v", err)
	}

	return nil
}

// SubidAlloc requests sysvisor-mgr to allocate a range of 'size' subuids and subgids.
func SubidAlloc(id string, size uint64) (uint32, uint32, error) {
	conn, err := connect()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SubidAllocReq{
		Id:   id,
		Size: size,
	}

	resp, err := ch.SubidAlloc(ctx, req)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to invoke SubidAlloc via grpc: %v", err)
	}

	return resp.Uid, resp.Gid, err
}

// ReqSupMounts requests sysvisor-mgr for supplementary mount configs for the container
// 'id' is the containers id
// 'rootfs' is the abs path to the container's rootfs
// 'uid' and 'gid' are the uid and gid of the container's root on the host
// 'shiftUids' indicates if sysvisor-runc is using uid-shifting for the container.
func ReqSupMounts(id string, rootfs string, uid, gid uint32, shiftUids bool) ([]*pb.Mount, error) {
	conn, err := connect()
	if err != nil {
		return []*pb.Mount{}, fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SupMountsReq{
		Id:        id,
		Rootfs:    rootfs,
		Uid:       uid,
		Gid:       gid,
		ShiftUids: shiftUids,
	}

	resp, err := ch.ReqSupMounts(ctx, req)
	if err != nil {
		return []*pb.Mount{}, fmt.Errorf("failed to invoke ReqSupMounts via grpc: %v", err)
	}

	return resp.GetMounts(), nil
}
