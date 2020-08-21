module grpc_code

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.31.0

replace golang.org/x/net => github.com/golang/net v0.0.0-20200813134508-3edf25e44fcc
