package sysvisorFsGrpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/nestybox/sysvisor/sysvisor-ipc/sysvisorFsGrpc/protobuf"
	"google.golang.org/grpc"
)

const sysvisorFsAddress = "localhost:50052"

// Container info passed by the client to the server across the grpc channel
type ContainerData struct {
	Id       string
	InitPid  int32
	Hostname string
	Ctime    time.Time
	UidFirst int32
	UidSize  int32
	GidFirst int32
	GidSize  int32
}

//
// Establishes grpc connection to sysvisor-fs' remote-end.
//
func connect() *grpc.ClientConn {

	// Set up a connection to the server.
	// TODO: Secure me through TLS.
	conn, err := grpc.Dial(sysvisorFsAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Sysvisorfs: %v", err)
		return nil
	}

	return conn
}

func containerDataToPbData(data *ContainerData) (*pb.ContainerData, error) {
	pbTime, err := ptypes.TimestampProto(data.Ctime)
	if err != nil {
		return nil, fmt.Errorf("time conversion error: %v", err)
	}

	return &pb.ContainerData{
		Id:      data.Id,
		InitPid: data.InitPid,
		Ctime:   pbTime,
	}, nil
}

func pbDatatoContainerData(pbdata *pb.ContainerData) (*ContainerData, error) {
	cTime, err := ptypes.Timestamp(pbdata.Ctime)
	if err != nil {
		return nil, fmt.Errorf("time conversion error: %v", err)
	}

	return &ContainerData{
		Id:      pbdata.Id,
		InitPid: pbdata.InitPid,
		Ctime:   cTime,
	}, nil
}

//
// Registers container creation in sysvisor-fs. Notice that this
// is a blocking call that can potentially have a minor impact
// on container's boot-up speed.
//
func SendContainerRegistration(data *ContainerData) (err error) {
	var pbData *pb.ContainerData

	// Set up sysvisorfs pipeline.
	conn := connect()
	if conn == nil {
		return fmt.Errorf("failed to connect with sysvisor-fs")
	}
	defer conn.Close()

	cntrChanIntf := pb.NewSysvisorStateChannelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pbData, err = containerDataToPbData(data)
	if err != nil {
		return fmt.Errorf("convertion to protobuf data failed: %v", err)
	}

	_, err = cntrChanIntf.ContainerRegistration(ctx, pbData)
	if err != nil {
		return fmt.Errorf("failed to register container with sysvisor-fs: %v", err)
	}

	return nil
}

//
// Unregisters container from Sysvisorfs.
//
func SendContainerUnregistration(data *ContainerData) (err error) {
	var pbData *pb.ContainerData

	// Set up sysvisorfs pipeline.
	conn := connect()
	if conn == nil {
		return fmt.Errorf("failed to connect with sysvisor-fs")
	}
	defer conn.Close()

	cntrChanIntf := pb.NewSysvisorStateChannelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pbData, err = containerDataToPbData(data)
	if err != nil {
		return fmt.Errorf("convertion to protobuf data failed: %v", err)
	}

	_, err = cntrChanIntf.ContainerUnregistration(ctx, pbData)
	if err != nil {
		return fmt.Errorf("failed to unregister container with sysvisor-fs: %v", err)
	}

	return nil
}

//
// Sends a container-update message to sysvisor-fs end. At this point, we are
// only utilizing this message for a particular case, update the container
// creation-time attribute, but this function can serve more general purposes
// in the future.
//
func SendContainerUpdate(data *ContainerData) (err error) {
	var pbData *pb.ContainerData

	// Set up sysvisorfs pipeline.
	conn := connect()
	if conn == nil {
		return fmt.Errorf("failed to connect with sysvisor-fs")
	}
	defer conn.Close()

	cntrChanIntf := pb.NewSysvisorStateChannelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pbData, err = containerDataToPbData(data)
	if err != nil {
		return fmt.Errorf("convertion to protobuf data failed: %v", err)
	}

	_, err = cntrChanIntf.ContainerUpdate(ctx, pbData)
	if err != nil {
		return fmt.Errorf("failed to send container-update message to ",
			"sysvisor-fs: %v", err)
	}

	return nil
}
