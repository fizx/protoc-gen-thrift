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

func map_string_Feature_ToProto(in map[string]*t.Feature) map[string]*example.Feature {
  if in == nil {
    return nil
  }
  out := make(map[string]*example.Feature)
  for k, v := range in {
    out[stringToProto(k)] = FeatureToProto(v)
  }
  return out
}

func map_string_Feature_ToThrift(in map[string]*example.Feature) map[string]*t.Feature {
  if in == nil {
    return nil
  }
  out := make(map[string]*t.Feature )
  for k, v := range in {
    out[stringToThrift(k)] = FeatureToThrift(v)
  }
  return out
}


func list_Entity_ToProto(in []*t.Entity) []*example.Entity {
  if in == nil {
    return nil
  }
  out := make([]*example.Entity, len(in))
  for i, e := range in {
    out[i] = EntityToProto(e)
  }
  return out
}

func list_Entity_ToThrift(in []*example.Entity) []*t.Entity {
  if in == nil {
    return nil
  }
  out := make([]*t.Entity, len(in))
  for i, e := range in {
    out[i] = EntityToThrift(e)
  }
  return out
}
//map[go_type:Entity proto_type:*example.Entity safe_thrift_type:Entity thrift_type:*t.Entity]

func list_string_ToProto(in []string) []string {
  if in == nil {
    return nil
  }
  out := make([]string, len(in))
  for i, e := range in {
    out[i] = stringToProto(e)
  }
  return out
}

func list_string_ToThrift(in []string) []string {
  if in == nil {
    return nil
  }
  out := make([]string, len(in))
  for i, e := range in {
    out[i] = stringToThrift(e)
  }
  return out
}
//map[go_type:string proto_type:string safe_thrift_type:string thrift_type:string]

func list_i64_ToProto(in []int64) []int64 {
  if in == nil {
    return nil
  }
  out := make([]int64, len(in))
  for i, e := range in {
    out[i] = i64ToProto(e)
  }
  return out
}

func list_i64_ToThrift(in []int64) []int64 {
  if in == nil {
    return nil
  }
  out := make([]int64, len(in))
  for i, e := range in {
    out[i] = i64ToThrift(e)
  }
  return out
}
//map[go_type:int64 proto_type:int64 safe_thrift_type:i64 thrift_type:int64]

  func RankingRequestToProto(in *t.RankingRequest) *example.RankingRequest {
    if in == nil {
      return nil
    }
    out := &example.RankingRequest{}
        out.Context = EntityToProto(in.Context)
    
        out.Candidates = list_Entity_ToProto(in.Candidates)
    
        out.Options = RequestOptionsToProto(in.Options)


    return out
  }
  func RankingRequestToThrift(in *example.RankingRequest) *t.RankingRequest {
    if in == nil {
      return nil
    }
    out := &t.RankingRequest{}
        out.Context = EntityToThrift(in.Context)
    
        out.Candidates = list_Entity_ToThrift(in.Candidates)
    
        out.Options = RequestOptionsToThrift(in.Options)


    return out
  }
  

  func EntityToProto(in *t.Entity) *example.Entity {
    if in == nil {
      return nil
    }
    out := &example.Entity{}
        out.Id = stringToProto(in.ID)
    
        out.Features = map_string_Feature_ToProto(in.Features)
    
        out.Score = doubleToProto(in.Score)


    return out
  }
  func EntityToThrift(in *example.Entity) *t.Entity {
    if in == nil {
      return nil
    }
    out := &t.Entity{}
        out.ID = stringToThrift(in.Id)
    
        out.Features = map_string_Feature_ToThrift(in.Features)
    
        out.Score = doubleToThrift(in.Score)


    return out
  }
  

  func FeatureToProto(in *t.Feature) *example.Feature {
    if in == nil {
      return nil
    }
    out := &example.Feature{}
    
    
    
    
    


        if in.ValueSelection == t.FeatureValue_AsString {
          out.Value = &example.Feature_AsString {
            AsString: stringToProto(in.AsString),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsInt {
          out.Value = &example.Feature_AsInt {
            AsInt: i64ToProto(in.AsInt),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsFloat {
          out.Value = &example.Feature_AsFloat {
            AsFloat: doubleToProto(in.AsFloat),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsBool {
          out.Value = &example.Feature_AsBool {
            AsBool: boolToProto(in.AsBool),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsStringArray {
          out.Value = &example.Feature_AsStringArray {
            AsStringArray: StringArrayToProto(in.AsStringArray),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsIntArray {
          out.Value = &example.Feature_AsIntArray {
            AsIntArray: IntArrayToProto(in.AsIntArray),
          }
        }

    return out
  }
  func FeatureToThrift(in *example.Feature) *t.Feature {
    if in == nil {
      return nil
    }
    out := &t.Feature{}
    
    
    
    
    

        if x, ok := in.GetValue().(*example.Feature_AsString); ok {
          out.ValueSelection = t.FeatureValue_AsString
		      out.AsString = stringToThrift(x.AsString)
	      }      
      
        if x, ok := in.GetValue().(*example.Feature_AsInt); ok {
          out.ValueSelection = t.FeatureValue_AsInt
		      out.AsInt = i64ToThrift(x.AsInt)
	      }      
      
        if x, ok := in.GetValue().(*example.Feature_AsFloat); ok {
          out.ValueSelection = t.FeatureValue_AsFloat
		      out.AsFloat = doubleToThrift(x.AsFloat)
	      }      
      
        if x, ok := in.GetValue().(*example.Feature_AsBool); ok {
          out.ValueSelection = t.FeatureValue_AsBool
		      out.AsBool = boolToThrift(x.AsBool)
	      }      
      
        if x, ok := in.GetValue().(*example.Feature_AsStringArray); ok {
          out.ValueSelection = t.FeatureValue_AsStringArray
		      out.AsStringArray = StringArrayToThrift(x.AsStringArray)
	      }      
      
        if x, ok := in.GetValue().(*example.Feature_AsIntArray); ok {
          out.ValueSelection = t.FeatureValue_AsIntArray
		      out.AsIntArray = IntArrayToThrift(x.AsIntArray)
	      }      

    return out
  }
  

  func StringArrayToProto(in *t.StringArray) *example.StringArray {
    if in == nil {
      return nil
    }
    out := &example.StringArray{}
        out.Value = list_string_ToProto(in.Value)


    return out
  }
  func StringArrayToThrift(in *example.StringArray) *t.StringArray {
    if in == nil {
      return nil
    }
    out := &t.StringArray{}
        out.Value = list_string_ToThrift(in.Value)


    return out
  }
  

  func IntArrayToProto(in *t.IntArray) *example.IntArray {
    if in == nil {
      return nil
    }
    out := &example.IntArray{}
        out.Value = list_i64_ToProto(in.Value)


    return out
  }
  func IntArrayToThrift(in *example.IntArray) *t.IntArray {
    if in == nil {
      return nil
    }
    out := &t.IntArray{}
        out.Value = list_i64_ToThrift(in.Value)


    return out
  }
  

  func RequestOptionsToProto(in *t.RequestOptions) *example.RequestOptions {
    if in == nil {
      return nil
    }
    out := &example.RequestOptions{}
        out.Limit = i32ToProto(in.Limit)


    return out
  }
  func RequestOptionsToThrift(in *example.RequestOptions) *t.RequestOptions {
    if in == nil {
      return nil
    }
    out := &t.RequestOptions{}
        out.Limit = i32ToThrift(in.Limit)


    return out
  }
  



  type RankingAdapterImpl struct {
    inner example.RankingServer
  }

  func (impl RankingAdapterImpl) IsHealthy(ctx context.Context) (bool, error) {
    _, err := impl.inner.IsHealthy(ctx, &empty.Empty{})
    return err == nil, err
  }
  
  
    func (impl RankingAdapterImpl) Rank(ctx context.Context, req *t.RankingRequest) (*t.RankingRequest, error) {
      out, err := impl.inner.Rank(ctx, RankingRequestToProto(req))
      return RankingRequestToThrift(out), err
    }

  func RankingAdapter(server example.RankingServer) t.Ranking {
    return &RankingAdapterImpl {
      inner: server,
    }
  }
