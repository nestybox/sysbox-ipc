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

//
// Sysbox-mgr config shared with other Sysbox components
//

type MgrConfig struct {
	AliasDns bool
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
