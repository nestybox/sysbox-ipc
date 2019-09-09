//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

// Server-side gRPC interface for the sysbox manager daemon

package sysboxMgrGrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	pb "github.com/nestybox/sysbox-ipc/sysboxMgrGrpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcSockAddr = "/run/sysbox/sysmgr.sock"

type ServerCallbacks struct {
	Register     func(id string) error
	Unregister   func(id string) error
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
	pb.RegisterSysboxMgrStateChannelServer(grpcServer, s)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *ServerStub) GetAddr() string {
	return grpcSockAddr
}

func (s *ServerStub) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	if req == nil {
		return &pb.RegisterResp{}, errors.New("invalid payload")
	}
	err := s.cb.Register(req.GetId())
	return &pb.RegisterResp{}, err
}

func (s *ServerStub) Unregister(ctx context.Context, req *pb.UnregisterReq) (*pb.UnregisterResp, error) {
	if req == nil {
		return &pb.UnregisterResp{}, errors.New("invalid payload")
	}
	err := s.cb.Unregister(req.GetId())
	return &pb.UnregisterResp{}, err
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
