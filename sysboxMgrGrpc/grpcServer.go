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
	"path"

	pb "github.com/nestybox/sysbox-ipc/sysboxMgrGrpc/protobuf"
	ipcLib "github.com/nestybox/sysbox-ipc/sysboxMgrLib"
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/opencontainers/runtime-spec/specs-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const sysMgrGrpcSockAddr = "/run/sysbox/sysmgr.sock"

type ServerCallbacks struct {
	Register       func(id string) error
	Unregister     func(id string) error
	SubidAlloc     func(id string, size uint64) (uint32, uint32, error)
	ReqMounts      func(id, rootfs string, uid, gid uint32, shiftUids bool, reqList []ipcLib.MountReqInfo) ([]specs.Mount, error)
	PrepMounts     func(id string, uid, gid uint32, shiftUids bool, prepList []ipcLib.MountPrepInfo) error
	ReqShiftfsMark func(id string, rootfs string, mounts []configs.ShiftfsMount) error
	Pause          func(id string) error
}

type ServerStub struct {
	cb *ServerCallbacks
}

func NewServerStub(cb *ServerCallbacks) *ServerStub {
	if cb == nil {
		return nil
	}

	if err := os.RemoveAll(sysMgrGrpcSockAddr); err != nil {
		return nil
	}

	if err := os.MkdirAll(path.Dir(sysMgrGrpcSockAddr), 0700); err != nil {
		return nil
	}

	return &ServerStub{
		cb: cb,
	}
}

func (s *ServerStub) Init() error {

	lis, err := net.Listen("unix", sysMgrGrpcSockAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if err := os.Chmod(sysMgrGrpcSockAddr, 0600); err != nil {
		return fmt.Errorf("failed to chmod %s: %v", sysMgrGrpcSockAddr, err)
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
	return sysMgrGrpcSockAddr
}

func (s *ServerStub) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	if req == nil {
		return &pb.RegisterResp{}, errors.New("invalid payload")
	}

	if err := s.cb.Register(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.RegisterResp{}, nil
}

func (s *ServerStub) Unregister(ctx context.Context, req *pb.UnregisterReq) (*pb.UnregisterResp, error) {
	if req == nil {
		return &pb.UnregisterResp{}, errors.New("invalid payload")
	}

	if err := s.cb.Unregister(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.UnregisterResp{}, nil
}

func (s *ServerStub) SubidAlloc(ctx context.Context, req *pb.SubidAllocReq) (*pb.SubidAllocResp, error) {
	if req == nil {
		return &pb.SubidAllocResp{}, errors.New("invalid payload")
	}

	uid, gid, err := s.cb.SubidAlloc(req.GetId(), req.GetSize())
	if err != nil {
		return nil, err
	}

	return &pb.SubidAllocResp{
		Uid: uid,
		Gid: gid,
	}, nil
}

func (s *ServerStub) ReqMounts(ctx context.Context, req *pb.MountReq) (*pb.MountResp, error) {
	if req == nil {
		return &pb.MountResp{}, errors.New("invalid payload")
	}

	// convert []*pb.MountReqInfo -> []ipcLib.MountReqInfo
	reqList := []ipcLib.MountReqInfo{}
	for _, pbInfo := range req.ReqList {
		info := ipcLib.MountReqInfo{
			Dest: pbInfo.GetDest(),
		}
		reqList = append(reqList, info)
	}

	mounts, err := s.cb.ReqMounts(req.GetId(), req.GetRootfs(), req.GetUid(), req.GetGid(), req.GetShiftUids(), reqList)
	if err != nil {
		return nil, err
	}

	// convert []*specs.Mount -> []*pb.Mount
	pbMounts := []*pb.Mount{}
	for _, m := range mounts {
		pbm := &pb.Mount{
			Source: m.Source,
			Dest:   m.Destination,
			Type:   m.Type,
			Opt:    m.Options,
		}
		pbMounts = append(pbMounts, pbm)
	}

	return &pb.MountResp{
		Mounts: pbMounts,
	}, nil
}

func (s *ServerStub) PrepMounts(ctx context.Context, req *pb.MountPrepReq) (*pb.MountPrepResp, error) {
	if req == nil {
		return &pb.MountPrepResp{}, errors.New("invalid payload")
	}

	// convert []*pb.MountPrepInfo -> []ipcLib.MountPrepInfo
	prepList := []ipcLib.MountPrepInfo{}
	for _, pbInfo := range req.PrepList {
		info := ipcLib.MountPrepInfo{
			Source:    pbInfo.GetSource(),
			Exclusive: pbInfo.GetExclusive(),
		}
		prepList = append(prepList, info)
	}

	if err := s.cb.PrepMounts(req.GetId(), req.GetUid(), req.GetGid(), req.GetShiftUids(), prepList); err != nil {
		return nil, err
	}

	return &pb.MountPrepResp{}, nil
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

	if err := s.cb.ReqShiftfsMark(req.GetId(), req.GetRootfs(), shiftfsMounts); err != nil {
		return nil, err
	}

	return &pb.ShiftfsMarkResp{}, nil
}

func (s *ServerStub) Pause(ctx context.Context, req *pb.PauseReq) (*pb.PauseResp, error) {
	if req == nil {
		return &pb.PauseResp{}, errors.New("invalid payload")
	}

	if err := s.cb.Pause(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.PauseResp{}, nil
}
