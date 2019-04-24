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
	SubidAlloc func(size uint64) (uint32, uint32, error)
	SubidFree  func(uid, gid uint32) error
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

func (s *ServerStub) SubidAlloc(ctx context.Context, req *pb.SubidAllocReq) (*pb.SubidAllocResp, error) {
	if req == nil {
		return &pb.SubidAllocResp{}, errors.New("invalid payload")
	}
	uid, gid, err := s.cb.SubidAlloc(req.GetSize())
	return &pb.SubidAllocResp{
		Uid: uid,
		Gid: gid,
	}, err
}

func (s *ServerStub) SubidFree(ctx context.Context, req *pb.SubidFreeReq) (*pb.SubidFreeResp, error) {
	if req == nil {
		return &pb.SubidFreeResp{}, errors.New("invalid payload")
	}
	err := s.cb.SubidFree(req.GetUid(), req.GetGid())
	return &pb.SubidFreeResp{}, err
}
