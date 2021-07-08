module github.com/nestybox/sysbox-ipc

go 1.13

require (
	github.com/golang/protobuf v1.4.1
	github.com/opencontainers/runc v0.0.0-00010101000000-000000000000
	github.com/opencontainers/runtime-spec v1.0.2
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f
	google.golang.org/grpc v1.27.0
)

replace github.com/opencontainers/runc => ../sysbox-runc

replace github.com/nestybox/sysbox-libs/libseccomp-golang => ../sysbox-libs/seccomp-golang

replace github.com/nestybox/sysbox-libs/formatter => ../sysbox-libs/formatter

replace github.com/godbus/dbus => github.com/godbus/dbus/v5 v5.0.3
