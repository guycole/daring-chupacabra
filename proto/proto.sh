#
rm chupacabra.pb.go
rm chupacabra_grpc.pb.go
#
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./chupacabra.proto
#