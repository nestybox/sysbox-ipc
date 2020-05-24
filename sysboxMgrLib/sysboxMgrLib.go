//
// Copyright: (C) 2019-2020 Nestybox Inc.  All rights reserved.
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
