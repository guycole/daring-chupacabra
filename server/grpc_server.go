// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"

	pb "github.com/guycole/daring-chupacabra/proto"
)

type ServerType struct {
	pb.UnimplementedChupacabraServer
}

func (ss *ServerType) EnqueueSubmit(ctx context.Context, in *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.EnqueueResponse{RequestStatus: 11, Token: "woot"}, nil
}
