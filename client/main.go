package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/guycole/daring-chupacabra/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const banner = "chupacapra-client 0.0"

func writeCommand(id, cmd string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cc := pb.NewChupacabraClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rr, err := cc.EnqueueSubmit(ctx, &pb.EnqueueRequest{Message: cmd, ClientId: id})
	if err != nil {
		log.Fatalf("could not write: %v", err)
	}

	log.Printf("Result: %s", rr.GetReceiptId())
}

func receiveTraffic(id string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cc := pb.NewChupacabraClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rr, err := cc.PollTest(ctx, &pb.PollRequest{ClientId: id})
	if err != nil {
		log.Fatalf("could not read: %v", err)
	}

	log.Printf("Result: %s", rr.GetClientId())
}

func main() {
	flag.Parse()

	log.Println(banner)

	clientId := uuid.NewString()
	log.Printf("client id:%s", clientId)

	runFlag := true
	for runFlag {
		fmt.Print("prompt>")

		var input string
		fmt.Scanln(&input)
		if input == "quit" {
			runFlag = false
			continue
		}

		writeCommand(clientId, input)
		receiveTraffic(clientId)
	}
}
