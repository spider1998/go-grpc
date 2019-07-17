codegen:customer.pb.go

customer.pb.go:customer.proto
	protoc  --go_out=plugins=grpc:. customer.proto

clean:
	-em -ef echo.pb.go
