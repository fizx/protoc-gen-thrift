package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/fizx/protoc-gen-thrift/example/generated/example"
	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/adapter"
	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/example/proto"
)

type server struct {
	RequestCounter int
}

func (srv *server) IsHealthy(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	srv.RequestCounter += 1
	return req, nil
}

func (srv *server) Rank(ctx context.Context, req *example.RankingRequest) (*example.RankingRequest, error) {
	srv.RequestCounter += 1
	return req, nil
}

func main() {
	service := &server{}
	adapt := adapter.RankingAdapter(service)
	processor := proto.NewRankingProcessor(adapt)
	fmt.Println("zomg i have a valid processor", processor)
}

var x example.RankingServer = &server{}
