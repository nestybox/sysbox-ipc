// Client-side gRPC interface for the sysvisor manager daemon

package sysvisorMgrGrpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/nestybox/sysvisor/sysvisor-protobuf/sysvisorMgrGrpc/protobuf"
	"google.golang.org/grpc"
)

// TODO: merge with grpcMaster; should be configurable ...
const sysvisorMgrAddr = "localhost:50053"

// connect establishes grpc connection to the sysvisorMgr daemon.
func connect() *grpc.ClientConn {

	// TODO: Secure me through TLS.
	conn, err := grpc.Dial(sysvisorMgrAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to SysvisorMgr: %v", err)
		return nil
	}

	return conn
}

// UidAlloc allocates an unused range of 'len' uids and gids
func UidAlloc(len uint32) (uint32, error) {
	conn := connect()
	if conn == nil {
		return 0, fmt.Errorf("failed to connect with sysvisorMgr")
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UidAllocReq{
		Len: len,
	}

	resp, err := ch.UidAlloc(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("failed to invoke uidAlloc via grpc: %v", err)
	}

	return resp.Uid, err
}

// UidFree releases a previously allocated id range; the given 'id' must
// be one previously returned by a successful call to UidAlloc()
// (otherwise this function returns a "not-found" error)
func UidFree(id uint32) error {

	conn := connect()
	if conn == nil {
		return fmt.Errorf("failed to connect with sysvisorMgr")
	}
	defer conn.Close()

	ch := pb.NewSysvisorMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UidFreeReq{
		Uid: id,
	}

	_, err := ch.UidFree(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke uidFree via grpc: %v", err)
	}

	return err
}
