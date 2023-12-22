package main

import (
	"context"
	"log"

	pb "github.com/guycole/daring-chupacabra/proto"
)

type serverType struct {
	pb.UnimplementedChupacabraServer
}

func (ss *serverType) EnqueueSubmit(ctx context.Context, in *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.EnqueueResponse{RequestStatus: 11, Token: "woot"}, nil
}
