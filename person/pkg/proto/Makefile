codegen:person.pb.go

person.pb.go:person.proto
	protoc  --go_out=plugins=grpc:. person.proto

clean:
	-rm -rf person.pb.go
