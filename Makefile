default:
	go generate .
	go build .

test: default
	rm -rf example/generated/thrift
	mkdir -p example/generated/thrift
	env PATH=.:$$PATH protoc -I/usr/local/include -Iexample --thrift_out=plugins=grpc,paths=source_relative:example/generated/thrift example/*.proto
	cd example && make test