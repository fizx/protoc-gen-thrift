package adapter

import (
  "context"
  "github.com/golang/protobuf/ptypes/empty"
  "github.com/fizx/protoc-gen-thrift/example/generated/example"
 t "github.com/fizx/protoc-gen-thrift/example/generated/thrift/example/example/proto"
)

func stringToProto(s string) string {
  return s
}
func stringToThrift(s string) string {
  return s
}
func i64ToProto(s int64) int64 {
  return s
}
func i64ToThrift(s int64) int64 {
  return s
}
func i32ToProto(s int32) int32 {
  return s
}
func i32ToThrift(s int32) int32 {
  return s
}
func boolToProto(s bool) bool {
  return s
}
func boolToThrift(s bool) bool {
  return s
}


func doubleToProto(s float64) float64 {
  return s
}
func doubleToThrift(s float64) float64 {
  return s
}



  func EmptyToProto(in *t.Empty) *example.Empty {
    if in == nil {
      return nil
    }
    out := &example.Empty{}


    return out
  }
  func EmptyToThrift(in *example.Empty) *t.Empty {
    if in == nil {
      return nil
    }
    out := &t.Empty{}


    return out
  }
  


