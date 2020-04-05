module github.com/fizx/protoc-gen-thrift

go 1.14

require (
	github.com/alexkappa/mustache v0.0.0-20191113130723-8bb9cfca2bfa
	github.com/golang/protobuf v1.3.5
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jwangsadinata/go-multimap v0.0.0-20190620162914-c29f3d7f33b6
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd // indirect
)

replace github.com/fizx/protoc-gen-thrift => ./

// replace github.com/fizx/protoc-gen-thrift/generated/assets => ./generated/assets
