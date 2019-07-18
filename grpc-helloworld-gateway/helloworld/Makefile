codegen:helloworld.pb.go
codegen:helloworld.pb.gw.go



helloworld.pb.go:helloworld.proto
	protoc -I/usr/local/include -I. \
    -I$$GOPATH/src \
    -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
    helloworld.proto

helloworld.pb.gw.go:helloworld.proto
	protoc -I/usr/local/include -I. \
    -I$$GOPATH/src  -I$$GOPATH/src/github.com/googleapis/googleapis/ \
    --grpc-gateway_out=logtostderr=true:. \
    helloworld.proto

clean:
	-rm -rf helloworld.pb.go
	-rm -rf helloworld.pb.gw.go

