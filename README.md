# protoc-gen-thrift  

protoc-gen-thrift takes a protobuf+GRPC IDL, and generates the corresponding Thrift IDL, as well as a golang adapter between them.  This allows you to write GRPC services, and also expose them as thrift endpoints for backwards compatibility with old thrift clients.

    # Usage:
    :$ go get -u github.com/fizx/protoc-gen-thrift 
    :$ git@github.com:fizx/protoc-gen-thrift.git
    :$ protoc -I/usr/local/include -I. --thrift_out=...    

There's a bit of ceremony involved in the invokation, so check out ./example/Makefile for an example of how to invoke it in a reasonable project.