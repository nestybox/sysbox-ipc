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

// Common definitions for grpc transfers with sysbox-mgr

package sysboxMgrLib

import (
	"github.com/nestybox/sysbox-libs/idShiftUtils"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Sysbox-mgr container registration info
type RegistrationInfo struct {
	Id          string
	Rootfs      string
	Userns      string
	Netns       string
	UidMappings []specs.LinuxIDMapping
	GidMappings []specs.LinuxIDMapping
}

// Sysbox-mgr container update info
type UpdateInfo struct {
	Id                 string
	Userns             string
	Netns              string
	UidMappings        []specs.LinuxIDMapping
	GidMappings        []specs.LinuxIDMapping
	RootfsUidShiftType idShiftUtils.IDShiftType
}

// Sysbox-mgr mandated container configs (passed from sysbox-mgr -> sysbox-runc)
type ContainerConfig struct {
	AliasDns                bool
	ShiftfsOk               bool
	ShiftfsOnOverlayfsOk    bool
	NoShiftfsOnFuse         bool
	IDMapMountOk            bool
	OverlayfsOnIDMapMountOk bool
	NoRootfsCloning         bool
	IgnoreSysfsChown        bool
	AllowTrustedXattr       bool
	HonorCaps               bool
	SyscontMode             bool
	Userns                  string
	UidMappings             []specs.LinuxIDMapping
	GidMappings             []specs.LinuxIDMapping
	FsuidMapFailOnErr       bool
	RootfsUidShiftType      idShiftUtils.IDShiftType
	RelaxedReadOnly         bool
}

//
// Mount requests from sysbox-runc to sysbox-mgr
//

type MountPrepInfo struct {
	Source    string
	Exclusive bool
}

type MntKind int

const (
	MntVarLibDocker MntKind = iota
	MntVarLibKubelet
	MntVarLibK0s
	MntVarLibRancherK3s
	MntVarLibRancherRke2
	MntVarLibBuildkit
	MntVarLibContainerdOvfs
	MntUsrSrcLinuxHdr
)

func (k MntKind) String() string {
	str := "unknown"

	switch k {
	case MntVarLibDocker:
		str = "var-lib-docker"
	case MntVarLibKubelet:
		str = "var-lib-kubelet"
	case MntVarLibK0s:
		str = "var-lib-k0s"
	case MntVarLibRancherK3s:
		str = "var-lib-rancher-k3s"
	case MntVarLibRancherRke2:
		str = "var-lib-rancher-rke2"
	case MntVarLibBuildkit:
		str = "var-lib-buildkit"
	case MntVarLibContainerdOvfs:
		str = "var-lib-containerd-ovfs"
	case MntUsrSrcLinuxHdr:
		str = "usr-src-linux-header"
	}

	return str
}

type MountReqInfo struct {
	Kind MntKind
	Dest string
}
