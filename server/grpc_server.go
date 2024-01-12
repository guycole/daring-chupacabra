// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	pb "github.com/guycole/daring-chupacabra/proto"
)

type CommandType struct {
	ExecuteTurn int
}

type ServerType struct {
	Eclectic *EclecticType
	SugarLog *zap.SugaredLogger

	pb.UnimplementedChupacabraServer
}

func (st *ServerType) EnqueueSubmit(ctx context.Context, in *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	st.SugarLog.Debug("enqueue submit")

	clientId := in.ClientId
	message := in.Message
	receiptId := uuid.NewString()

	ent := EventNodeType{Action: parseAction, ClientID: clientId, RawCommand: message, ReceiptID: receiptId}
	st.Eclectic.insertNodeNextTurn(&ent)

	return &pb.EnqueueResponse{ClientId: clientId, ReceiptId: receiptId}, nil
}

func (st *ServerType) PollTest(ctx context.Context, in *pb.PollRequest) (*pb.PollResponse, error) {
	st.SugarLog.Debug("poll test")

	clientId := in.ClientId

	return &pb.PollResponse{ClientId: clientId, Responses: []*pb.PollResponse_ResponseTraffic{}}, nil
}
