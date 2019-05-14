// Server-side gRPC interface for the sysvisor manager daemon

package sysvisorMgrGrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	pb "github.com/nestybox/sysvisor/sysvisor-ipc/sysvisorMgrGrpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcSockAddr = "/run/sysvisor/sysmgr.sock"

type ServerCallbacks struct {
	SubidAlloc   func(id string, size uint64) (uint32, uint32, error)
	SubidFree    func(id string) error
	ReqSupMounts func(id string, rootfs string, uid, gid uint32, shiftUids bool) ([]*pb.Mount, error)
	RelSupMounts func(id string) error
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
	if err := os.RemoveAll(grpcSockAddr); err != nil {
		return err
	}

	lis, err := net.Listen("unix", grpcSockAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if err := os.Chmod(grpcSockAddr, 0600); err != nil {
		return fmt.Errorf("failed to chmod %s: %v", grpcSockAddr, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSysvisorMgrStateChannelServer(grpcServer, s)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *ServerStub) GetAddr() string {
	return grpcSockAddr
}

func (s *ServerStub) SubidAlloc(ctx context.Context, req *pb.SubidAllocReq) (*pb.SubidAllocResp, error) {
	if req == nil {
		return &pb.SubidAllocResp{}, errors.New("invalid payload")
	}
	uid, gid, err := s.cb.SubidAlloc(req.GetId(), req.GetSize())
	return &pb.SubidAllocResp{
		Uid: uid,
		Gid: gid,
	}, err
}

func (s *ServerStub) SubidFree(ctx context.Context, req *pb.SubidFreeReq) (*pb.SubidFreeResp, error) {
	if req == nil {
		return &pb.SubidFreeResp{}, errors.New("invalid payload")
	}
	err := s.cb.SubidFree(req.GetId())
	return &pb.SubidFreeResp{}, err
}

func (s *ServerStub) ReqSupMounts(ctx context.Context, req *pb.SupMountsReq) (*pb.SupMountsReqResp, error) {
	if req == nil {
		return &pb.SupMountsReqResp{}, errors.New("invalid payload")
	}

	mounts, err := s.cb.ReqSupMounts(req.GetId(), req.GetRootfs(), req.GetUid(), req.GetGid(), req.GetShiftUids())
	if err != nil {
		return &pb.SupMountsReqResp{}, err
	}

	return &pb.SupMountsReqResp{
		Mounts: mounts,
	}, nil

}

func (s *ServerStub) RelSupMounts(ctx context.Context, req *pb.SupMountsRel) (*pb.SupMountsRelResp, error) {
	if req == nil {
		return &pb.SupMountsRelResp{}, errors.New("invalid payload")
	}
	err := s.cb.RelSupMounts(req.GetId())
	return &pb.SupMountsRelResp{}, err
}
