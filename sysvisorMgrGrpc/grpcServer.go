// Server-side gRPC interface for the sysvisor manager daemon

package sysvisorMgrGrpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/nestybox/sysvisor/sysvisor-protobuf/sysvisorMgrGrpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// TODO: get rid of this; use unix domain socket instead
	// TODO: merge this with sysvisorMgrAddr; same for sysvisorFsGrpc ...
	port = ":50053"
)

type ServerCallbacks struct {
	UidAlloc func(len uint32) (uint32, error)
	UidFree  func(id uint32) error
}

type ServerStub struct {
	cb *ServerCallbacks
}

func NewServerStub(cb *ServerCallbacks) *ServerStub {
	if cb == nil {
		return nil
	}
	return &ServerStub{
		cb: cb,
	}
}

func (s *ServerStub) Init() error {

	// TODO: Change me to unix domain socket instead: more secure/efficient.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSysvisorMgrStateChannelServer(grpcServer, s)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

// Generates a uid/gid allocation request
func (s *ServerStub) UidAlloc(ctx context.Context, req *pb.UidAllocReq) (*pb.UidAllocResp, error) {
	if req == nil {
		return &pb.UidAllocResp{}, errors.New("invalid payload")
	}
	uid, err := s.cb.UidAlloc(req.GetLen())
	return &pb.UidAllocResp{Uid: uid}, err
}

// Generates a uid/gid freeing request
func (s *ServerStub) UidFree(ctx context.Context, req *pb.UidFreeReq) (*pb.UidFreeResp, error) {
	if req == nil {
		return &pb.UidFreeResp{}, errors.New("invalid payload")
	}
	err := s.cb.UidFree(req.GetUid())
	return &pb.UidFreeResp{}, err
}
