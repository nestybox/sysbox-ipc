package sysvisorFsGrpc

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/nestybox/sysvisor-ipc/sysvisorFsGrpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//
// File dealing with all the logic related to Sysvisorfs' external-communication
// (ipc) logic.
//


const (
      sysvisorFsPort = ":50052"
      sysvisorFsAddress = "localhost" + sysvisorFsPort
)


const (
	Unknown MessageType = iota
	ContainerRegisterMessage
	ContainerUnregisterMessage
	ContainerUpdateMessage
	MaxSupportedMessage
)

var messageTypeStrings = [...]string{
	"Unknown",
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

	newServer := &Server{
		Ctx:       ctx,
		Callbacks: make(map[MessageType]Callback),
	}

	for ctype, cval := range *cb {
		newServer.Callbacks[ctype] = cval
	}

	return newServer
}

func (s *Server) Init() {

	// TODO: Change me to unix-socket instead: more secure/efficient.
	lis, err := net.Listen("tcp", sysvisorFsPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initializing grpc server
	grpcServer := grpc.NewServer()

	pb.RegisterSysvisorStateChannelServer(grpcServer, s)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// TODO: To be implemented in the future if needed.
func (s *Server) CallbackRegister(c *Callback) {

}

// TODO: To be implemented in the future if needed.
func (s *Server) CallbackUnregister(c *Callback) {

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

	if data == nil {
		return &pb.Response{Success: false},
			errors.New("gRPC server: invalid msg payload received")
	}

	// Verify incoming message-request is supported.

	// Obtain the associated callback matching this incoming request.
	cb, ok := s.Callbacks[mtype]
	if !ok {
		return &pb.Response{Success: false},
			errors.New("gRPC server: no callback registered for incoming msg")
	}

	// Transform received payload to a grpc/protobuf-agnostic message.
	cont, err := pbDatatoContainerData(data)
	if err != nil {
		return &pb.Response{Success: false}, err
	}

        err = (cb)(s.Ctx, cont)
	if err != nil {
		return &pb.Response{Success: false},
			errors.New("gRPC server: unexpected response from client endpoint")
	}

	return &pb.Response{Success: true}, nil
}
