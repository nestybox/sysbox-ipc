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

// Client-side gRPC interface for the sysbox manager daemon

package sysboxMgrGrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/nestybox/sysbox-ipc/sysboxMgrGrpc/sysboxMgrProtobuf"
	ipcLib "github.com/nestybox/sysbox-ipc/sysboxMgrLib"
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/opencontainers/runtime-spec/specs-go"
	"google.golang.org/grpc"
)

func unixConnect(addr string, t time.Duration) (net.Conn, error) {
	unixAddr, err := net.ResolveUnixAddr("unix", sysMgrGrpcSockAddr)
	conn, err := net.DialUnix("unix", nil, unixAddr)
	return conn, err
}

// connect establishes grpc connection to the sysbox-mgr daemon.
func connect() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(sysMgrGrpcSockAddr, grpc.WithInsecure(), grpc.WithDialer(unixConnect))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Registers a container with sysbox-mgr
func Register(id string) (*ipcLib.MgrConfig, error) {
	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &pb.RegisterReq{
		Id: id,
	}

	resp, err := ch.Register(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke Register via grpc: %v", err)
	}

	config := &ipcLib.MgrConfig{
		AliasDns: resp.MgrConfig.GetAliasDns(),
	}

	return config, nil
}

// Unregisters a container with sysbox-mgr
func Unregister(id string) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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

// SubidAlloc requests sysbox-mgr to allocate a range of 'size' subuids and subgids.
func SubidAlloc(id string, size uint64) (uint32, uint32, error) {
	conn, err := connect()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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

// ReqMounts requests the sysbox-mgr to setup sys container special mounts
func ReqMounts(id, rootfs string, uid, gid uint32, shiftUids bool, reqList []ipcLib.MountReqInfo) ([]specs.Mount, error) {

	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	// We don't use context timeout for this API because the time it takes to
	// setup the mounts can be large, in particular for sys containers that come
	// preloaded with heavy inner images and in machines where the load is high.
	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Convert []ipcLib.MountReqInfo -> []*pb.MountReqInfo
	pbReqList := []*pb.MountReqInfo{}
	for _, info := range reqList {
		pbInfo := &pb.MountReqInfo{
			Kind: uint32(info.Kind),
			Dest: info.Dest,
		}
		pbReqList = append(pbReqList, pbInfo)
	}

	req := &pb.MountReq{
		Id:        id,
		Rootfs:    rootfs,
		Uid:       uid,
		Gid:       gid,
		ShiftUids: shiftUids,
		ReqList:   pbReqList,
	}

	resp, err := ch.ReqMounts(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke ReqMounts via grpc: %v", err)
	}

	// Convert []*pb.Mount -> []specs.Mount
	specMounts := []specs.Mount{}
	for _, m := range resp.Mounts {
		specm := specs.Mount{
			Source:      m.GetSource(),
			Destination: m.GetDest(),
			Type:        m.GetType(),
			Options:     m.GetOpt(),
		}
		specMounts = append(specMounts, specm)
	}

	return specMounts, nil
}

// PrepMounts requests sysbox-mgr to prepare a mount source for use by a sys container.
func PrepMounts(id string, uid, gid uint32, shiftUids bool, prepList []ipcLib.MountPrepInfo) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Convert []ipcLib.MountPrepInfo -> []*pb.MountPrepInfo
	pbPrepList := []*pb.MountPrepInfo{}
	for _, info := range prepList {
		pbInfo := &pb.MountPrepInfo{
			Source:    info.Source,
			Exclusive: info.Exclusive,
		}
		pbPrepList = append(pbPrepList, pbInfo)
	}

	req := &pb.MountPrepReq{
		Id:        id,
		Uid:       uid,
		Gid:       gid,
		ShiftUids: shiftUids,
		PrepList:  pbPrepList,
	}

	_, err = ch.PrepMounts(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke PrepMounts via grpc: %v", err)
	}

	return nil
}

// ReqShiftfsMark requests sysbox-mgr to perform shiftfs marking on the container's
// rootfs and the given list of other mountpoints.
func ReqShiftfsMark(id string, rootfs string, mounts []configs.ShiftfsMount) error {

	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// convert configs.ShiftfsMount to grpc ShiftfsMark
	shiftfsMarks := []*pb.ShiftfsMark{}
	for _, m := range mounts {
		sm := &pb.ShiftfsMark{
			Source:   m.Source,
			Readonly: m.Readonly,
		}
		shiftfsMarks = append(shiftfsMarks, sm)
	}

	req := &pb.ShiftfsMarkReq{
		Id:           id,
		Rootfs:       rootfs,
		ShiftfsMarks: shiftfsMarks,
	}

	_, err = ch.ReqShiftfsMark(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke ReqShiftfsMark via grpc: %v", err)
	}

	return nil
}

// Pause notifies the sysbox-mgr that the container has been paused.
// 'id' is the containers id
func Pause(id string) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &pb.PauseReq{
		Id: id,
	}

	_, err = ch.Pause(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke Pause via grpc: %v", err)
	}

	return nil
}
