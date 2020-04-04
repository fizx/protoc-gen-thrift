default:
	go build .

test: default
	rm -rf generated
	mkdir -p generated
	env PATH=.:$$PATH protoc -I/usr/local/include -I. --thrift_out=plugins=grpc,paths=source_relative:generated/ *.proto