package main

import (
	"io/ioutil"
	"os"

	"github.com/alexkappa/mustache"
	"github.com/fizx/protoc-gen-thrift/generated/assets"
	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

//go:generate go run gen/main.go
func main() {
	fs := assets.Assets
	input, _ := ioutil.ReadAll(os.Stdin)
	req := &plugin_go.CodeGeneratorRequest{}
	proto.Unmarshal(input, req)
	m := mustache.New()
	for _, file := range req.GetProtoFile() {
		file.GetPackage()
	}

	rsp := &plugin_go.CodeGeneratorResponse{}
	bytes, _ := proto.Marshal(rsp)
	os.Stdout.Write(bytes)
}
