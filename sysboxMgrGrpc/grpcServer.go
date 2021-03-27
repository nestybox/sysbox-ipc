//
// Copyright 2019-2020 Nestybox, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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

	pb "github.com/nestybox/sysbox-ipc/sysboxMgrGrpc/sysboxMgrProtobuf"
	ipcLib "github.com/nestybox/sysbox-ipc/sysboxMgrLib"
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/opencontainers/runtime-spec/specs-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const sysMgrGrpcSockAddr = "/run/sysbox/sysmgr.sock"

type ServerCallbacks struct {
	Register       func(id, userns, netns string) (*ipcLib.MgrConfig, error)
	Unregister     func(id string) error
	SubidAlloc     func(id string, size uint64) (uint32, uint32, error)
	ReqMounts      func(id, rootfs string, uid, gid uint32, shiftUids bool, reqList []ipcLib.MountReqInfo) ([]specs.Mount, error)
	PrepMounts     func(id string, uid, gid uint32, shiftUids bool, prepList []ipcLib.MountPrepInfo) error
	ReqShiftfsMark func(id string, rootfs string, mounts []configs.ShiftfsMount) error
	ReqFsState     func(id string, rootfs string) ([]configs.FsEntry, error)
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

	config, err := s.cb.Register(req.GetId(), req.GetUserns(), req.GetNetns())
	if err != nil {
		return nil, err
	}

	mgrConfig := pb.MgrConfig{
		AliasDns: config.AliasDns,
		Userns:   config.Userns,
	}

	resp := &pb.RegisterResp{
		MgrConfig: &mgrConfig,
	}

	return resp, nil
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
			Kind: ipcLib.MntKind(pbInfo.GetKind()),
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

func (s *ServerStub) ReqFsState(
	ctx context.Context,
	req *pb.FsStateReq) (*pb.FsStateResp, error) {

	if req == nil {
		return &pb.FsStateResp{}, errors.New("invalid payload")
	}

	fsState, err := s.cb.ReqFsState(req.GetId(), req.GetRootfs())
	if err != nil {
		return nil, err
	}

	// convert []configs.FsEntry -> []*pb.FsEntry
	pbFsEntries := []*pb.FsEntry{}
	for _, e := range fsState {
		pbe := &pb.FsEntry{
			Kind: uint32(e.GetKind()),
			Path: e.GetPath(),
			Mode: uint32(e.GetMode()),
			Dst:  e.GetDest(),
		}
		pbFsEntries = append(pbFsEntries, pbe)
	}

	return &pb.FsStateResp{FsEntries: pbFsEntries}, nil
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
