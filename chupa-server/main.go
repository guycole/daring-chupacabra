package main

import (
	"flag"
	"log"
	// pb "github.com/guycole/daring-chupacabra/proto/chupacabra.pb.go"
)

var (
	port = flag.Int("port", 50051, "server port")
)

/*
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}
*/

/*
// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
*/

const banner = "chupa-server 0.0"

// Remove the duplicate package declaration
// package main

var (
	port = flag.Int("port", 50051, "server port")
)

func main() {
	flag.Parse()

	log.Println(banner)

	sr := &SubmitRequest{}
	log.Printf("sr: %v", sr)

	/*
	   lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	   	if err != nil {
	   		log.Fatalf("failed to listen: %v", err)
	   	}

	   s := grpc.NewServer()
	   pb.RegisterGreeterServer(s, &server{})
	   log.Printf("server listening at %v", lis.Addr())

	   	if err := s.Serve(lis); err != nil {
	   		log.Fatalf("failed to serve: %v", err)
	   	}
	*/
}
