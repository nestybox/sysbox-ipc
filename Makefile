#
# sysbox-ipc Makefile
#

.PHONY: clean sysbox-ipc sysbox-ipc-debug sysbox-ipc-static

GO := go

SYSIPC_GRPC_FS_DIR  := sysboxFsGrpc
SYSIPC_GRPC_MGR_DIR := sysboxMgrGrpc
SYSIPC_UNIX_DIR     := unix

SYSIPC_SRC := $(shell find $(SYSIPC_GRPC_FS_DIR) 2>&1 | egrep ".*\.(go|proto)")
SYSIPC_SRC += $(shell find $(SYSIPC_GRPC_MGR_DIR) 2>&1 | egrep ".*\.(go|proto)")
SYSIPC_SRC += $(shell find $(SYSIPC_UNIX_DIR) 2>&1 | egrep ".*\.(go|proto)")

LDFLAGS := '-X main.version=${VERSION} -X main.commitId=${COMMIT_ID} \
			-X "main.builtAt=${BUILT_AT}" -X main.builtBy=${BUILT_BY}'

sysbox-ipc: $(SYSIPC_SRC) sysipc-grpc-fs-proto sysipc-grpc-mgr-proto
	$(GO) build -ldflags ${LDFLAGS} ./$(SYSIPC_GRPC_FS_DIR)
	$(GO) build -ldflags ${LDFLAGS} ./$(SYSIPC_GRPC_MGR_DIR)
	$(GO) build -ldflags ${LDFLAGS} ./$(SYSIPC_UNIX_DIR)

sysbox-ipc-debug: $(SYSIPC_SRC) sysipc-grpc-fs-proto sysipc-grpc-mgr-proto
	$(GO) build -gcflags="all=-N -l" ./...

sysbox-ipc-static: $(SYSIPC_SRC) sysipc-grpc-fs-proto sysipc-grpc-mgr-proto
	CGO_ENABLED=1 $(GO) build -tags "netgo osusergo static_build" \
		-installsuffix netgo -ldflags "-w -extldflags -static" ./...

sysipc-grpc-fs-proto:
#	@cd $(SYSIPC_GRPC_FS_DIR)/protobuf && make
	$(MAKE) -C $(SYSIPC_GRPC_FS_DIR)/protobuf

sysipc-grpc-mgr-proto:
#	@cd $(SYSIPC_GRPC_MGR_DIR)/protobuf && make
	$(MAKE) -C $(SYSIPC_GRPC_MGR_DIR)/protobuf

clean:
	$(MAKE) -C $(SYSIPC_GRPC_FS_DIR)/protobuf clean
	$(MAKE) -C $(SYSIPC_GRPC_MGR_DIR)/protobuf clean
