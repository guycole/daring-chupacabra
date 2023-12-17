package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/guycole/daring-chupacabra/proto"
)

var (
	port = flag.Int("port", 50051, "server port")
)

type server struct {
	pb.UnimplementedChupacabraServer
}

func (ss *server) EnqueueSubmit(ctx context.Context, in *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.EnqueueResponse{RequestStatus: 11, Token: "woot"}, nil
}

const banner = "chupacapra-server 0.0"

func main() {
	flag.Parse()

	log.Println(banner)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ss := grpc.NewServer()
	pb.RegisterChupacabraServer(ss, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := ss.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
