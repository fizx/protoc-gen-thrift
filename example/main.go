package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/fizx/protoc-gen-thrift/example/generated/example"
	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/adapter"
	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/example/proto"
)

type server struct{}

func (server) IsHealthy(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return req, nil
}

func (server) Rank(ctx context.Context, req *example.RankingRequest) (*example.RankingRequest, error) {
	return req, nil
}

func main() {
	service := &server{}
	adapt := adapter.RankingAdapter(service)
	processor := proto.NewRankingProcessor(adapt)
	fmt.Println("zomg i have a valid processor", processor)
}

var x example.RankingServer = &server{}
