module github.com/nestybox/sysbox-ipc

go 1.21

require (
	github.com/golang/protobuf v1.5.4
	github.com/nestybox/sysbox-libs/formatter v0.0.0-00010101000000-000000000000
	github.com/nestybox/sysbox-libs/idShiftUtils v0.0.0-00010101000000-000000000000
	github.com/nestybox/sysbox-libs/shiftfs v0.0.0-00010101000000-000000000000
	github.com/opencontainers/runc v1.1.4
	github.com/opencontainers/runtime-spec v1.1.1-0.20230823135140-4fec88fd00a4
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/sys v0.20.0
	google.golang.org/grpc v1.63.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/coreos/go-systemd/v22 v22.1.0 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/docker/docker v26.0.0+incompatible // indirect
	github.com/godbus/dbus/v5 v5.0.3 // indirect
	github.com/joshlf/go-acl v0.0.0-20200411065538-eae00ae38531 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/nestybox/sysbox-libs/linuxUtils v0.0.0-00010101000000-000000000000 // indirect
	github.com/nestybox/sysbox-libs/mount v0.0.0-00010101000000-000000000000 // indirect
	github.com/nestybox/sysbox-libs/utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	gopkg.in/hlandau/service.v1 v1.0.7 // indirect
)

replace (
	github.com/godbus/dbus => github.com/godbus/dbus/v5 v5.0.3
	github.com/nestybox/sysbox-ipc => ./
	github.com/nestybox/sysbox-libs/capability => ../sysbox-libs/capability
	github.com/nestybox/sysbox-libs/dockerUtils => ../sysbox-libs/dockerUtils
	github.com/nestybox/sysbox-libs/formatter => ../sysbox-libs/formatter
	github.com/nestybox/sysbox-libs/idShiftUtils => ../sysbox-libs/idShiftUtils
	github.com/nestybox/sysbox-libs/libseccomp-golang => ../sysbox-libs/libseccomp-golang
	github.com/nestybox/sysbox-libs/linuxUtils => ../sysbox-libs/linuxUtils
	github.com/nestybox/sysbox-libs/mount => ../sysbox-libs/mount
	github.com/nestybox/sysbox-libs/shiftfs => ../sysbox-libs/shiftfs
	github.com/nestybox/sysbox-libs/utils => ../sysbox-libs/utils
	github.com/opencontainers/runc => ../sysbox-runc
)
