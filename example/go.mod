module github.com/fizx/protoc-gen-thrift/example

go 1.14

replace github.com/fizx/protoc-gen-thrift/example => ./

require (
	github.com/apache/thrift v0.13.0
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.28.0
)
