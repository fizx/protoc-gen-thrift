# protoc-gen-thrift  

protoc-gen-thrift takes a protobuf+GRPC IDL, and generates the corresponding Thrift IDL, as well as a golang adapter between them.  This allows you to write GRPC services, and also expose them as thrift endpoints for backwards compatibility with old thrift clients.

    # Usage:
    :$ go get -u github.com/fizx/protoc-gen-thrift 
    :$ git@github.com:fizx/protoc-gen-thrift.git
    :$ protoc -I/usr/local/include -I. --thrift_out=...    

There's a bit of ceremony involved in the invokation, so check out ./example/Makefile for an example of how to invoke it in a reasonable project.

Code's still pretty ugly, and if you do some odd things in protobuf, I'm sure it'll mildly break.  Also, package selection is pretty broken.  All this is easy to fix, but hey, super-alpha stuff.

Also, thrift doesn't support oneofs, so we add an extra field to thrift requests in order to specify which field we mean:

    &proto.Feature{
        ValueSelection: proto.FeatureValue_AsString,
        AsString:       "hello",
	}