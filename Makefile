#
# sysbox-ipc Makefile
#

.PHONY: clean sysbox-ipc sysipc-grpc-fs-proto sysipc-grpc-mgr-proto validate list-packages

GO := go

SYSIPC_GRPC_FS_DIR  := sysboxFsGrpc
SYSIPC_GRPC_MGR_DIR := sysboxMgrGrpc

sysbox-ipc: sysipc-grpc-fs-proto sysipc-grpc-mgr-proto

sysipc-grpc-fs-proto:
	$(MAKE) -C $(SYSIPC_GRPC_FS_DIR)/sysboxFsProtobuf

sysipc-grpc-mgr-proto:
	$(MAKE) -C $(SYSIPC_GRPC_MGR_DIR)/sysboxMgrProtobuf

validate:
	script/validate-gofmt

listpackages:
	@echo $(allpackages)

clean:
	$(MAKE) -C $(SYSIPC_GRPC_FS_DIR)/sysboxFsProtobuf clean
	$(MAKE) -C $(SYSIPC_GRPC_MGR_DIR)/sysboxMgrProtobuf clean

# memoize allpackages, so that it's executed only once and only if used
_allpackages = $(shell $(GO) list ./... | grep -v vendor)
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)
