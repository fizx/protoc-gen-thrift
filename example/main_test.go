package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/adapter"
	"github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/example/proto"
)

func TestAdapter(t *testing.T) {
	service := &server{}
	adapt := adapter.RankingAdapter(service)
	processor := proto.NewRankingProcessor(adapt)
	req := proto.NewRankingRequest()
	rsp, _ := adapt.Rank(context.Background(), req)
	assert.Equal(t, req, rsp)
	fmt.Println("zomg i have a valid processor", processor)
}

func TestAdapterAgain(t *testing.T) {
	service := &server{}
	adapt := adapter.RankingAdapter(service)
	processor := proto.NewRankingProcessor(adapt)
	f := make(map[string]*proto.Feature)
	f["foo"] = &proto.Feature{
		ValueSelection: proto.FeatureValue_AsString,
		AsString:       "hello",
	}
	req := &proto.RankingRequest{
		Context: &proto.Entity{
			ID:       "hello",
			Features: f,
			Score:    9.8,
		},
	}
	rsp, _ := adapt.Rank(context.Background(), req)
	assert.Equal(t, req, rsp)
	assert.Equal(t, 1, service.RequestCounter)
	fmt.Println("zomg i have a valid processor", processor)
}
