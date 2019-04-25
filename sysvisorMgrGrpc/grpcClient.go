// Client-side gRPC interface for the sysvisor manager daemon

package sysvisorMgrGrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/nestybox/sysvisor/sysvisor-protobuf/sysvisorMgrGrpc/protobuf"
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

// SubidAlloc requests sysvisor-mgr to allocate a range of 'size' subuids and subgids
func SubidAlloc(size uint64) (uint32, uint32, error) {
	conn, err := connect()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SubidAllocReq{
		Size: size,
	}

	resp, err := ch.SubidAlloc(ctx, req)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to invoke SubidAlloc via grpc: %v", err)
	}

	return resp.Uid, resp.Gid, err
}

// SubidFree releases previously allocated uid and gid ranges; the given uid and gid must
// be from a previous call to SubidAlloc() (otherwise this function returns a "not-found"
// error)
func SubidFree(uid, gid uint32) error {

	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysvisor-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SubidFreeReq{
		Uid: uid,
		Gid: gid,
	}

	_, err = ch.SubidFree(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke subidFree via grpc: %v", err)
	}

	return err
}
