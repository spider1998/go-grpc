codegen:customer.pb.go

customer.pb.go:customer.proto
	protoc  --go_out=plugins=grpc:. customer.proto

clean:
	-rm -rf echo.pb.go
