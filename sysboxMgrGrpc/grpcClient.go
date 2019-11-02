//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

// Client-side gRPC interface for the sysbox manager daemon

package sysboxMgrGrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/nestybox/sysbox-ipc/sysboxMgrGrpc/protobuf"
	"github.com/opencontainers/runc/libcontainer/configs"
	"google.golang.org/grpc"
)

func unixConnect(addr string, t time.Duration) (net.Conn, error) {
	unixAddr, err := net.ResolveUnixAddr("unix", grpcSockAddr)
	conn, err := net.DialUnix("unix", nil, unixAddr)
	return conn, err
}

// connect establishes grpc connection to the sysbox-mgr daemon.
func connect() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grpcSockAddr, grpc.WithInsecure(), grpc.WithDialer(unixConnect))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Registers a container with sysbox-mgr
func Register(id string) error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &pb.RegisterReq{
		Id: id,
	}

	_, err = ch.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to invoke Register via grpc: %v", err)
	}

	return nil
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
// `mode` indicates the allocation mode to be used; if empty, Sysbox's default allocation
// mode is used.
func SubidAlloc(id string, size uint64, mode string) (uint32, uint32, error) {
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
		Mode: mode,
	}

	resp, err := ch.SubidAlloc(ctx, req)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to invoke SubidAlloc via grpc: %v", err)
	}

	return resp.Uid, resp.Gid, err
}

// ReqSupMounts requests sysbox-mgr for supplementary mount configs for the container
// 'id' is the containers id
// 'rootfs' is the abs path to the container's rootfs
// 'uid' and 'gid' are the uid and gid of the container's root on the host
// 'shiftUids' indicates if sysbox-runc is using uid-shifting for the container.
func ReqSupMounts(id string, rootfs string, uid, gid uint32, shiftUids bool) ([]*pb.Mount, error) {
	conn, err := connect()
	if err != nil {
		return []*pb.Mount{}, fmt.Errorf("failed to connect with sysbox-mgr: %v", err)
	}
	defer conn.Close()

	ch := pb.NewSysboxMgrStateChannelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &pb.SupMountsReq{
		Id:        id,
		Rootfs:    rootfs,
		Uid:       uid,
		Gid:       gid,
		ShiftUids: shiftUids,
	}

	resp, err := ch.ReqSupMounts(ctx, req)
	if err != nil {
		return []*pb.Mount{}, fmt.Errorf("failed to invoke ReqSupMounts via grpc: %v", err)
	}

	return resp.GetMounts(), nil
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
