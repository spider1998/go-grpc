codegen:counter.pb.go



counter.pb.go:counter.proto
	protoc --go_out=plugins=grpc:. counter.proto

clean:
	-em -ef counter.pb.go
