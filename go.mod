module github.com/nestybox/sysbox-ipc

go 1.13

require (
	github.com/golang/protobuf v1.3.1
	github.com/opencontainers/runc v0.0.0-00010101000000-000000000000
	github.com/opencontainers/runtime-spec v1.0.2-0.20191218002532-bab266ed033b
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449
	google.golang.org/grpc v1.21.0
)

replace github.com/opencontainers/runc => ../sysbox-runc

replace github.com/nestybox/libseccomp-golang => ../lib/seccomp-golang
