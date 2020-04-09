package adapter

import (
  "context"
  "github.com/golang/protobuf/ptypes/empty"
 p "github.com/fizx/protoc-gen-thrift/example/generated/example"
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

func map_string_Feature_ToProto(in map[string]*t.Feature) map[string]*p.Feature {
  if in == nil {
    return nil
  }
  out := make(map[string]*p.Feature)
  for k, v := range in {
    out[stringToProto(k)] = FeatureToProto(v)
  }
  return out
}

func map_string_Feature_ToThrift(in map[string]*p.Feature) map[string]*t.Feature {
  if in == nil {
    return nil
  }
  out := make(map[string]*t.Feature )
  for k, v := range in {
    out[stringToThrift(k)] = FeatureToThrift(v)
  }
  return out
}


func list_Entity_ToProto(in []*t.Entity) []*p.Entity {
  if in == nil {
    return nil
  }
  out := make([]*p.Entity, len(in))
  for i, e := range in {
    out[i] = EntityToProto(e)
  }
  return out
}

func list_Entity_ToThrift(in []*p.Entity) []*t.Entity {
  if in == nil {
    return nil
  }
  out := make([]*t.Entity, len(in))
  for i, e := range in {
    out[i] = EntityToThrift(e)
  }
  return out
}
//map[go_type:Entity proto_type:*p.Entity safe_thrift_type:Entity thrift_type:*t.Entity]

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

  func RankingRequestToProto(in *t.RankingRequest) *p.RankingRequest {
    if in == nil {
      return nil
    }
    out := &p.RankingRequest{}
        out.Context = EntityToProto(in.Context)
    
        out.Candidates = list_Entity_ToProto(in.Candidates)
    
        out.Options = RequestOptionsToProto(in.Options)


    return out
  }
  func RankingRequestToThrift(in *p.RankingRequest) *t.RankingRequest {
    if in == nil {
      return nil
    }
    out := &t.RankingRequest{}
        out.Context = EntityToThrift(in.Context)
    
        out.Candidates = list_Entity_ToThrift(in.Candidates)
    
        out.Options = RequestOptionsToThrift(in.Options)


    return out
  }
  

  func EntityToProto(in *t.Entity) *p.Entity {
    if in == nil {
      return nil
    }
    out := &p.Entity{}
        out.Id = stringToProto(in.ID)
    
        out.Features = map_string_Feature_ToProto(in.Features)
    
        out.Score = doubleToProto(in.Score)


    return out
  }
  func EntityToThrift(in *p.Entity) *t.Entity {
    if in == nil {
      return nil
    }
    out := &t.Entity{}
        out.ID = stringToThrift(in.Id)
    
        out.Features = map_string_Feature_ToThrift(in.Features)
    
        out.Score = doubleToThrift(in.Score)


    return out
  }
  

  func FeatureToProto(in *t.Feature) *p.Feature {
    if in == nil {
      return nil
    }
    out := &p.Feature{}
    
    
    
    
    


        if in.ValueSelection == t.FeatureValue_AsString {
          out.Value = &p.Feature_AsString {
            AsString: stringToProto(in.AsString),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsInt {
          out.Value = &p.Feature_AsInt {
            AsInt: i64ToProto(in.AsInt),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsFloat {
          out.Value = &p.Feature_AsFloat {
            AsFloat: doubleToProto(in.AsFloat),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsBool {
          out.Value = &p.Feature_AsBool {
            AsBool: boolToProto(in.AsBool),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsStringArray {
          out.Value = &p.Feature_AsStringArray {
            AsStringArray: StringArrayToProto(in.AsStringArray),
          }
        }
      
        if in.ValueSelection == t.FeatureValue_AsIntArray {
          out.Value = &p.Feature_AsIntArray {
            AsIntArray: IntArrayToProto(in.AsIntArray),
          }
        }

    return out
  }
  func FeatureToThrift(in *p.Feature) *t.Feature {
    if in == nil {
      return nil
    }
    out := &t.Feature{}
    
    
    
    
    

        if x, ok := in.GetValue().(*p.Feature_AsString); ok {
          out.ValueSelection = t.FeatureValue_AsString
		      out.AsString = stringToThrift(x.AsString)
	      }      
      
        if x, ok := in.GetValue().(*p.Feature_AsInt); ok {
          out.ValueSelection = t.FeatureValue_AsInt
		      out.AsInt = i64ToThrift(x.AsInt)
	      }      
      
        if x, ok := in.GetValue().(*p.Feature_AsFloat); ok {
          out.ValueSelection = t.FeatureValue_AsFloat
		      out.AsFloat = doubleToThrift(x.AsFloat)
	      }      
      
        if x, ok := in.GetValue().(*p.Feature_AsBool); ok {
          out.ValueSelection = t.FeatureValue_AsBool
		      out.AsBool = boolToThrift(x.AsBool)
	      }      
      
        if x, ok := in.GetValue().(*p.Feature_AsStringArray); ok {
          out.ValueSelection = t.FeatureValue_AsStringArray
		      out.AsStringArray = StringArrayToThrift(x.AsStringArray)
	      }      
      
        if x, ok := in.GetValue().(*p.Feature_AsIntArray); ok {
          out.ValueSelection = t.FeatureValue_AsIntArray
		      out.AsIntArray = IntArrayToThrift(x.AsIntArray)
	      }      

    return out
  }
  

  func StringArrayToProto(in *t.StringArray) *p.StringArray {
    if in == nil {
      return nil
    }
    out := &p.StringArray{}
        out.Value = list_string_ToProto(in.Value)


    return out
  }
  func StringArrayToThrift(in *p.StringArray) *t.StringArray {
    if in == nil {
      return nil
    }
    out := &t.StringArray{}
        out.Value = list_string_ToThrift(in.Value)


    return out
  }
  

  func IntArrayToProto(in *t.IntArray) *p.IntArray {
    if in == nil {
      return nil
    }
    out := &p.IntArray{}
        out.Value = list_i64_ToProto(in.Value)


    return out
  }
  func IntArrayToThrift(in *p.IntArray) *t.IntArray {
    if in == nil {
      return nil
    }
    out := &t.IntArray{}
        out.Value = list_i64_ToThrift(in.Value)


    return out
  }
  

  func RequestOptionsToProto(in *t.RequestOptions) *p.RequestOptions {
    if in == nil {
      return nil
    }
    out := &p.RequestOptions{}
        out.Limit = i32ToProto(in.Limit)


    return out
  }
  func RequestOptionsToThrift(in *p.RequestOptions) *t.RequestOptions {
    if in == nil {
      return nil
    }
    out := &t.RequestOptions{}
        out.Limit = i32ToThrift(in.Limit)


    return out
  }
  



  type RankingAdapterImpl struct {
    inner p.RankingServer
  }

  func (impl RankingAdapterImpl) IsHealthy(ctx context.Context) (bool, error) {
    _, err := impl.inner.IsHealthy(ctx, &empty.Empty{})
    return err == nil, err
  }
  
  
    func (impl RankingAdapterImpl) Rank(ctx context.Context, req *t.RankingRequest) (*t.RankingRequest, error) {
      out, err := impl.inner.Rank(ctx, RankingRequestToProto(req))
      return RankingRequestToThrift(out), err
    }

  func RankingAdapter(server p.RankingServer) t.Ranking {
    return &RankingAdapterImpl {
      inner: server,
    }
  }
