module github.com/nestybox/sysbox-ipc

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/nestybox/sysbox-libs/formatter v0.0.0-00010101000000-000000000000
	github.com/nestybox/sysbox-libs/idShiftUtils v0.0.0-00010101000000-000000000000
	github.com/opencontainers/runc v1.0.0-rc9.0.20210126000000-2be806d1391d
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	google.golang.org/grpc v1.34.1
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/nestybox/sysbox-ipc => ./
	github.com/opencontainers/runc => ../sysbox-runc
	github.com/nestybox/sysbox-libs/libseccomp-golang => ../sysbox-libs/libseccomp-golang
	github.com/nestybox/sysbox-libs/formatter => ../sysbox-libs/formatter
	github.com/nestybox/sysbox-libs/capability => ../sysbox-libs/capability
	github.com/nestybox/sysbox-libs/utils => ../sysbox-libs/utils
	github.com/nestybox/sysbox-libs/dockerUtils => ../sysbox-libs/dockerUtils
	github.com/nestybox/sysbox-libs/idShiftUtils => ../sysbox-libs/idShiftUtils
	github.com/godbus/dbus => github.com/godbus/dbus/v5 v5.0.3
)
