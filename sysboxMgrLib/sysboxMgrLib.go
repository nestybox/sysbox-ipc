//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

// Common definitions for gprc transfers with sysbox-mgr

package sysboxMgrLib

type MountPrepInfo struct {
	Source    string
	Exclusive bool
}

type MountReqInfo struct {
	Dest string
}
