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
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/opencontainers/runtime-spec/specs-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcSockAddr = "/run/sysbox/sysmgr.sock"

type ServerCallbacks struct {
	Register             func(id string) error
	Unregister           func(id string) error
	SubidAlloc           func(id string, size uint64) (uint32, uint32, error)
	ReqDockerStoreMount  func(id string, rootfs string, uid, gid uint32, shiftUids bool) (*specs.Mount, error)
	PrepDockerStoreMount func(id string, path string, uid, gid uint32, shiftUids bool) error
	ReqShiftfsMark       func(id string, rootfs string, mounts []configs.ShiftfsMount) error
	Pause                func(id string) error
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

func (s *ServerStub) ReqDockerStoreMount(ctx context.Context, req *pb.DsMountReq) (*pb.DsMountResp, error) {
	if req == nil {
		return &pb.DsMountResp{}, errors.New("invalid payload")
	}

	m, err := s.cb.ReqDockerStoreMount(req.GetId(), req.GetRootfs(), req.GetUid(), req.GetGid(), req.GetShiftUids())
	if err != nil {
		return &pb.DsMountResp{}, err
	}

	pbMount := &pb.Mount{
		Source: m.Source,
		Dest:   m.Destination,
		Type:   m.Type,
		Opt:    m.Options,
	}

	return &pb.DsMountResp{
		Mount: pbMount,
	}, nil
}

func (s *ServerStub) PrepDockerStoreMount(ctx context.Context, req *pb.DsMountPrepReq) (*pb.DsMountPrepResp, error) {
	if req == nil {
		return &pb.DsMountPrepResp{}, errors.New("invalid payload")
	}

	err := s.cb.PrepDockerStoreMount(req.GetId(), req.GetPath(), req.GetUid(), req.GetGid(), req.GetShiftUids())
	return &pb.DsMountPrepResp{}, err
}

func (s *ServerStub) ReqShiftfsMark(ctx context.Context, req *pb.ShiftfsMarkReq) (*pb.ShiftfsMarkResp, error) {
	if req == nil {
		return &pb.ShiftfsMarkResp{}, errors.New("invalid payload")
	}

	// Convert pb.ShiftfsMark to configs.ShiftfsMount
	shiftfsMounts := []configs.ShiftfsMount{}
	for _, m := range req.GetShiftfsMarks() {
		sm := configs.ShiftfsMount{
			Source:   m.Source,
			Readonly: m.Readonly,
		}
		shiftfsMounts = append(shiftfsMounts, sm)
	}

	err := s.cb.ReqShiftfsMark(req.GetId(), req.GetRootfs(), shiftfsMounts)
	return &pb.ShiftfsMarkResp{}, err
}

func (s *ServerStub) Pause(ctx context.Context, req *pb.PauseReq) (*pb.PauseResp, error) {
	if req == nil {
		return &pb.PauseResp{}, errors.New("invalid payload")
	}
	err := s.cb.Pause(req.GetId())
	return &pb.PauseResp{}, err
}
