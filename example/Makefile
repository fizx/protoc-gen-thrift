default: gen-go
	go build .

test: default
	go test .

gen-go:
	rm -rf generated/thrift
	mkdir -p generated/thrift
	env PATH=..:$$PATH protoc -I/usr/local/include -I. --thrift_out=protopkg=github.com/fizx/protoc-gen-thrift/example/generated/example,thriftpkg=github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/example/proto:generated/thrift *.proto
	mkdir -p generated/example generated/thrift/example
	protoc -I/usr/local/include -I. --go_out=plugins=grpc,paths=source_relative:generated/example *.proto
	thrift -out generated/thrift/example --gen go:thrift_import=github.com/apache/thrift/lib/go/thrift,package_prefix=github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/ generated/thrift/baseplate.thrift
	thrift -out generated/thrift/example --gen go:thrift_import=github.com/apache/thrift/lib/go/thrift,package_prefix=github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/ generated/thrift/example.proto.thrift