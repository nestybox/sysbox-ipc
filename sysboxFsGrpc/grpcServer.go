//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

package sysboxFsGrpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path"

	pb "github.com/nestybox/sysbox-ipc/sysboxFsGrpc/sysboxFsProtobuf"

	"google.golang.org/grpc"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	grpcStatus "google.golang.org/grpc/status"
)

//
// File dealing with all the logic related to sysbox-fs' external-communication
// (ipc) logic.
//

const sysFsGrpcSockAddr = "/run/sysbox/sysfs.sock"

const (
	Unknown MessageType = iota
	ContainerPreRegisterMessage
	ContainerRegisterMessage
	ContainerUnregisterMessage
	ContainerUpdateMessage
	MaxSupportedMessage
)

var messageTypeStrings = [...]string{
	"Unknown",
	"ContainerPreRegister",
	"ContainerRegister",
	"ContainerUnregister",
	"ContainerUpdate",
}

type MessageType uint16

type Callback func(client interface{}, c *ContainerData) error

type CallbacksMap = map[MessageType]Callback

type Server struct {
	Ctx       interface{}
	Callbacks CallbacksMap
}

func NewServer(ctx interface{}, cb *CallbacksMap) *Server {

	if cb == nil {
		return nil
	}

	if err := os.RemoveAll(sysFsGrpcSockAddr); err != nil {
		return nil
	}

	if err := os.MkdirAll(path.Dir(sysFsGrpcSockAddr), 0700); err != nil {
		return nil
	}

	newServer := &Server{
		Ctx:       ctx,
		Callbacks: make(map[MessageType]Callback),
	}

	for ctype, cval := range *cb {
		newServer.Callbacks[ctype] = cval
	}

	return newServer
}

func (s *Server) Init() error {

	lis, err := net.Listen("unix", sysFsGrpcSockAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if err := os.Chmod(sysFsGrpcSockAddr, 0600); err != nil {
		return fmt.Errorf("failed to chmod %s: %v", sysFsGrpcSockAddr, err)
	}

	// Initializing grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterSysboxStateChannelServer(grpcServer, s)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

// TODO: To be implemented in the future if needed.
func (s *Server) CallbackRegister(c *Callback) {

}

// TODO: To be implemented in the future if needed.
func (s *Server) CallbackUnregister(c *Callback) {

}

func (s *Server) ContainerPreRegistration(
	ctx context.Context, data *pb.ContainerData) (*pb.Response, error) {

	return s.executeCallback(ContainerPreRegisterMessage, data)
}

func (s *Server) ContainerRegistration(
	ctx context.Context, data *pb.ContainerData) (*pb.Response, error) {

	return s.executeCallback(ContainerRegisterMessage, data)
}

func (s *Server) ContainerUnregistration(
	ctx context.Context, data *pb.ContainerData) (*pb.Response, error) {

	return s.executeCallback(ContainerUnregisterMessage, data)
}

func (s *Server) ContainerUpdate(
	ctx context.Context, data *pb.ContainerData) (*pb.Response, error) {

	return s.executeCallback(ContainerUpdateMessage, data)
}

func (s *Server) executeCallback(mtype MessageType,
	data *pb.ContainerData) (*pb.Response, error) {

	// Sanity-check data field here to avoid doing it in server backend.
	if data == nil {
		return &pb.Response{Success: false},
			grpcStatus.Error(grpcCodes.InvalidArgument, "Invalid data field")
	}

	// Obtain the associated callback matching this incoming request.
	cb, ok := s.Callbacks[mtype]
	if !ok {
		return &pb.Response{Success: false},
			grpcStatus.Errorf(
				grpcCodes.Unimplemented,
				"Method type %v not implemented",
				mtype,
			)
	}

	// Transform received payload to a grpc/protobuf-agnostic message.
	cont, err := pbDatatoContainerData(data)
	if err != nil {
		return &pb.Response{Success: false}, err
	}

	err = (cb)(s.Ctx, cont)
	if err != nil {
		return &pb.Response{Success: false}, err
	}

	return &pb.Response{Success: true}, nil
}
