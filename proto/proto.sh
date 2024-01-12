#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
rm chupacabra.pb.go
rm chupacabra_grpc.pb.go
#
# gsc@lingling:19>/usr/local/bin/protoc --version
# libprotoc 25.1
#
# brew install protoc-gen-go-grpc
#
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./chupacabra.proto
#
