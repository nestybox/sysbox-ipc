.PHONY: clean

RUNC_FS_SRC := protobuf.proto
RUNC_FS_PB := protobuf.pb.go

.DEFAULT: $(RUNC_FS_PB)

$(RUNC_FS_PB): $(RUNC_FS_SRC)
	protoc -I . -I /usr/local/include/ protobuf.proto --go_out=plugins=grpc:.

clean:
	rm -f $(RUNC_FS_PB)
