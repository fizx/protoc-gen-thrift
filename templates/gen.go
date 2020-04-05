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

{{#maps}}
func {{safe_name}}ToProto(in map[{{left_thrift}}]{{right_thrift}}) map[{{left_proto}}]{{right_proto}} {
  if in == nil {
    return nil
  }
  out := make(map[{{left_proto}}]{{right_proto}})
  for k, v := range in {
    out[{{left_type}}ToProto(k)] = {{right_type}}ToProto(v)
  }
  return out
}

func {{safe_name}}ToThrift(in map[{{left_proto}}]{{right_proto}}) map[{{left_thrift}}]{{right_thrift}} {
  if in == nil {
    return nil
  }
  out := make(map[{{left_thrift}}]{{right_thrift}} )
  for k, v := range in {
    out[{{left_type}}ToThrift(k)] = {{right_type}}ToThrift(v)
  }
  return out
}

{{/maps}}

{{#lists}}
func list_{{safe_thrift_type}}_ToProto(in []{{thrift_type}}) []{{proto_type}} {
  if in == nil {
    return nil
  }
  out := make([]{{proto_type}}, len(in))
  for i, e := range in {
    out[i] = {{safe_thrift_type}}ToProto(e)
  }
  return out
}

func list_{{safe_thrift_type}}_ToThrift(in []{{proto_type}}) []{{thrift_type}} {
  if in == nil {
    return nil
  }
  out := make([]{{thrift_type}}, len(in))
  for i, e := range in {
    out[i] = {{safe_thrift_type}}ToThrift(e)
  }
  return out
}
//{{.}}
{{/lists}}

{{#messageType}}
  func {{name}}ToProto(in *t.{{name}}) *example.{{name}} {
    if in == nil {
      return nil
    }
    out := &example.{{name}}{}
    {{#field}}
      {{^is_one_of}}
        out.{{camel_name}} = {{safe_thrift_type}}ToProto(in.{{special_camel_name}})
      {{/is_one_of}}
    {{/field}}

    {{#oneofs}}

      {{#value}}
        if in.{{key}}Selection == t.{{name}}{{key}}_{{inner_name}} {
          out.{{key}} = &example.{{name}}_{{inner_name}} {
            {{inner_name}}: {{inner_type}}ToProto(in.{{inner_name}}),
          }
        }
      {{/value}}
    {{/oneofs}}

    return out
  }
  func {{name}}ToThrift(in *example.{{name}}) *t.{{name}} {
    if in == nil {
      return nil
    }
    out := &t.{{name}}{}
    {{#field}}
      {{^is_one_of}}
        out.{{special_camel_name}} = {{safe_thrift_type}}ToThrift(in.{{prefix}}{{camel_name}})
      {{/is_one_of}}
    {{/field}}

    {{#oneofs}}
      {{#value}}
        if x, ok := in.Get{{key}}().(*example.{{name}}_{{inner_name}}); ok {
          out.{{key}}Selection = t.{{name}}{{key}}_{{inner_name}}
		      out.{{inner_name}} = {{inner_type}}ToThrift(x.{{inner_name}})
	      }      
      {{/value}}
    {{/oneofs}}

    return out
  }
  
{{/messageType}}


{{#service}}

  type {{name}}AdapterImpl struct {
    inner example.{{name}}Server
  }

  func (impl {{name}}AdapterImpl) IsHealthy(ctx context.Context) (bool, error) {
    _, err := impl.inner.IsHealthy(ctx, &empty.Empty{})
    return err == nil, err
  }
  
  {{#method}}
    {{^health_check}}
    func (impl {{service_name}}AdapterImpl) {{name}}(ctx context.Context, req *t.{{inputType}}) (*t.{{outputType}}, error) {
      out, err := impl.inner.{{name}}(ctx, {{inputType}}ToProto(req))
      return {{outputType}}ToThrift(out), err
    }
    {{/health_check}}
  {{/method}}

  func {{name}}Adapter(server example.{{name}}Server) t.Ranking {
    return &{{name}}AdapterImpl {
      inner: server,
    }
  }
{{/service}}