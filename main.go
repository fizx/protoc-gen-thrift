package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/alexkappa/mustache"
	"github.com/fizx/protoc-gen-thrift/generated/assets"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/iancoleman/strcase"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

func eachify(m *slicemultimap.MultiMap) []interface{} {
	a := make([]interface{}, 0)
	for _, k := range m.KeySet() {
		inner := make(map[interface{}]interface{})
		inner["key"] = k
		inner["value"], _ = m.Get(k)
		a = append(a, inner)
	}
	return a
}

func keysOf(m map[string]bool) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getTemplate(fs http.FileSystem, path string) *mustache.Template {
	m := mustache.New()
	file, _ := fs.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	m.ParseString(string(bytes))
	return m
}

var golangTypeMap = map[string]string{
	"STRING": "string",
	"INT64":  "int64",
	"BOOL":   "bool",
	"DOUBLE": "float64",
	"INT32":  "int32",
}

var thriftTypeMap = map[string]string{
	"STRING": "string",
	"INT64":  "i64",
	"BOOL":   "bool",
	"DOUBLE": "double",
	"INT32":  "i32",
}

//json array
func ja(x interface{}) []interface{} {
	return x.([]interface{})
}

//json array of maps
func jam(x interface{}) []map[string]interface{} {
	if x == nil {
		return make([]map[string]interface{}, 0)
	}
	a := x.([]interface{})
	out := make([]map[string]interface{}, len(a))
	for i, e := range a {
		out[i] = e.(map[string]interface{})
	}
	return out
}

//json map
func jm(x interface{}) map[string]interface{} {
	return x.(map[string]interface{})
}

var special = map[string]string{
	"Id": "ID",
}

func camel(s string) string {
	out := strcase.ToCamel(s)
	if special[out] != "" {
		out = special[out]
	}
	return out
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

func goType(x interface{}) string {
	key := x.(string)
	if strings.HasPrefix(key, "TYPE_") {
		key = key[5:]
	}
	ttype := golangTypeMap[key]
	if ttype == "" {
		ttype = key
	}
	return ttype
}

func typeFromMap(x interface{}) string {
	key := x.(string)[5:]
	ttype := thriftTypeMap[key]
	if ttype == "" {
		ttype = key
	}
	return ttype
}

func parseParams(p string) map[string]interface{} {
	m := make(map[string]interface{})
	list := strings.Split(p, ",")
	for _, e := range list {
		ee := strings.Split(e, "=")
		m[ee[0]] = ee[1]
	}
	return m
}

func merge(target *map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		(*target)[k] = v
	}
}

//go:generate go run gen/main.go
func main() {
	for k, v := range thriftTypeMap {
		golangTypeMap[v] = golangTypeMap[k]
	}
	fs := assets.Assets
	input, _ := ioutil.ReadAll(os.Stdin)
	req := &plugin_go.CodeGeneratorRequest{}
	rsp := &plugin_go.CodeGeneratorResponse{}
	_ = proto.Unmarshal(input, req)
	m := &jsonpb.Marshaler{}

	jsonBytes, _ := m.MarshalToString(req)
	data := make(map[string]interface{})
	json.Unmarshal([]byte(jsonBytes), &data)
	fmt.Fprintln(os.Stderr, parseParams(req.GetParameter()))
	for _, file := range jam(data["protoFile"]) {
		merge(&file, parseParams(req.GetParameter()))
		lists := make(map[string]bool)
		maps := make(map[string]bool)
		file["package"] = strings.Split(file["name"].(string), ".")[0]
		for _, service := range jam(file["service"]) {
			for _, method := range jam(service["method"]) {
				n := method["name"].(string)
				method["service_name"] = service["name"]
				if strings.HasSuffix(n, "ealthy") {
					method["health_check"] = true
				} else {
					method["inputType"] = method["inputType"].(string)[1:]
					method["outputType"] = method["outputType"].(string)[1:]
				}
			}
		}
		for _, msg := range jam(file["messageType"]) {
			od := jam(msg["oneofDecl"])
			oneofs := make([]string, len(od))
			components := slicemultimap.New()
			for i, oneof := range od {
				oneofs[i] = oneof["name"].(string)
			}
			msg["thrift_type"] = "struct"
			for _, field := range jam(msg["field"]) {
				field["camel_name"] = strcase.ToCamel(field["name"].(string))
				field["special_camel_name"] = camel(field["name"].(string))
				if field["type"] == "TYPE_MESSAGE" {
					typeNames := strings.Split(field["typeName"].(string), ".")
					if len(typeNames) == 2 {
						field["thrift_type"] = typeNames[1]
					} else {
						nname := typeNames[2]
						for _, nested := range jam(msg["nestedType"]) {
							if nested["name"] == nname {
								k := "UNKNOWN"
								v := "UNKNOWN"
								for _, nf := range jam(nested["field"]) {
									if nf["name"] == "key" {
										if nf["type"] == "TYPE_MESSAGE" {
											k = nf["typeName"].(string)[1:]
										} else {
											k = typeFromMap(nf["type"])
										}
									}
									if nf["name"] == "value" {
										if nf["type"] == "TYPE_MESSAGE" {
											v = nf["typeName"].(string)[1:]
										} else {
											v = typeFromMap(nf["type"])
										}
									}
								}
								field["thrift_type"] = "map<" + k + ", " + v + ">"
								maps[k+","+v] = true
								field["is_map"] = true
							}
						}
					}

				} else {
					ttype := typeFromMap(field["type"])
					field["thrift_type"] = ttype
				}
				if field["label"] == "LABEL_REPEATED" && field["is_map"] == nil {
					lists[field["thrift_type"].(string)] = true
					field["thrift_type"] = "list<" + field["thrift_type"].(string) + ">"
				}
				reg := regexp.MustCompile("[\\W]+")
				field["safe_thrift_type"] = reg.ReplaceAllString(field["thrift_type"].(string), "_")

				if field["oneofIndex"] != nil {
					both := make(map[string]interface{})
					both["inner_name"] = field["camel_name"]
					both["inner_type"] = field["thrift_type"]
					components.Put(camel(oneofs[int32(field["oneofIndex"].(float64))]), both)
					field["is_one_of"] = true
				}
			}

			msg["oneofs"] = eachify(components)
		}
		ll := make([]map[string]interface{}, 0)
		for k, _ := range lists {
			out := make(map[string]interface{})
			if isLower(k) {
				out["thrift_type"] = goType(k)
				out["proto_type"] = goType(k)
			} else {
				out["thrift_type"] = "*t." + k
				out["proto_type"] = "*p." + k
			}
			reg := regexp.MustCompile("[\\W]+")
			out["safe_thrift_type"] = reg.ReplaceAllString(k, "_")
			out["go_type"] = goType(k)
			ll = append(ll, out)
		}
		ml := make([]map[string]interface{}, 0)
		for k, _ := range maps {
			out := make(map[string]interface{})
			x := strings.Split(k, ",")
			l := x[0]
			r := x[1]
			out["name"] = k
			out["safe_name"] = "map_" + l + "_" + r + "_"
			out["left_type"] = l
			if isLower(l) {
				out["left_thrift"] = l
				out["left_proto"] = l
			} else {
				out["left_thrift"] = "*t." + l
				out["left_proto"] = "*p." + l
			}
			out["right_type"] = r
			if isLower(r) {
				out["right_thrift"] = r
				out["right_proto"] = r
			} else {
				out["right_thrift"] = "*t." + r
				out["right_proto"] = "*p." + r
			}
			ml = append(ml, out)
		}
		file["maps"] = ml
		file["lists"] = ll

		gotmpl := getTemplate(fs, "/gen.go")
		out, _ := gotmpl.RenderString(file)
		newName := "adapter/" + file["name"].(string) + ".go"
		gofile := plugin_go.CodeGeneratorResponse_File{
			Name:    &newName,
			Content: &out,
		}
		rsp.File = append(rsp.File, &gofile)

		ttmpl := getTemplate(fs, "/gen.thrift")
		tout, _ := ttmpl.RenderString(file)
		newTName := file["name"].(string) + ".thrift"
		tfile := plugin_go.CodeGeneratorResponse_File{
			Name:    &newTName,
			Content: &tout,
		}
		rsp.File = append(rsp.File, &tfile)

	}
	f, _ := fs.Open("/baseplate.thrift")
	bp, _ := ioutil.ReadAll(f)
	bps := string(bp)
	newBName := "baseplate.thrift"
	bfile := plugin_go.CodeGeneratorResponse_File{
		Name:    &newBName,
		Content: &bps,
	}
	rsp.File = append(rsp.File, &bfile)

	bytes, _ := proto.Marshal(rsp)
	os.Stdout.Write(bytes)
}
