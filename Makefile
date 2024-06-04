CH_PLUGIN=bin/protoc-gen-ch-schema
GO_PLUGIN=bin/protoc-gen-go
PROTOC_GEN_GO_PKG=google.golang.org/protobuf/cmd/protoc-gen-go
GLOG_PKG=github.com/golang/glog
PROTO_SRC=ch_table.proto ch_field.proto
PROTO_GENFILES=protos/ch_table.pb.go protos/ch_field.pb.go
PROTO_PKG=google.golang.org/protobuf
PKGMAP=Mgoogle/protobuf/descriptor.proto=$(PROTO_PKG)/types/descriptorpb
EXAMPLES_PROTO=examples/test_merge_tree.proto
EXAMPLES_REPLACING_PROTO=examples/test_replacing_merge_tree.proto
EXAMPLES_SUMMING_PROTO=examples/test_summing_merge_tree.proto

install: $(CH_PLUGIN)

$(CH_PLUGIN): $(PROTO_GENFILES) goprotobuf glog
	go build -o $@

$(PROTO_GENFILES): $(PROTO_SRC) $(GO_PLUGIN)
	protoc -I. --plugin=$(GO_PLUGIN) --go_out=$(PKGMAP):protos --go_opt=paths=source_relative $(PROTO_SRC)

goprotobuf:
	go get $(PROTO_PKG)

glog:
	go get $(GLOG_PKG)

$(GO_PLUGIN):
	go build -o $@ $(PROTOC_GEN_GO_PKG)

test: $(PROTO_SRC)
	go test

distclean clean:
	go clean
	rm -f $(GO_PLUGIN) $(CH_PLUGIN)

realclean: distclean
	rm -f $(PROTO_GENFILES)

examples: $(CH_PLUGIN)
	protoc -I. --plugin=$(CH_PLUGIN) --ch-schema_out=examples/schema $(EXAMPLES_PROTO)
	protoc -I. --plugin=$(CH_PLUGIN) --ch-schema_out=examples/schema $(EXAMPLES_REPLACING_PROTO)
	protoc -I. --plugin=$(CH_PLUGIN) --ch-schema_out=examples/schema $(EXAMPLES_SUMMING_PROTO)

.PHONY: goprotobuf glog
