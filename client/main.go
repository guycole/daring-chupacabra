package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/guycole/daring-chupacabra/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const banner = "chupacapra-client 0.0"

func writeCommand(cmd string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cc := pb.NewChupacabraClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rr, err := cc.EnqueueSubmit(ctx, &pb.EnqueueRequest{ExecuteTurn: -1, Message: "woot", Owner: "guy"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Result: %s", rr.GetToken())
}

func main() {
	flag.Parse()

	log.Println(banner)

	runFlag := true
	for runFlag {
		fmt.Print("prompt>")

		var input string
		fmt.Scanln(&input)

		fmt.Println(input)

		writeCommand(input)
	}

	/*
	   conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	   	if err != nil {
	   		log.Fatalf("did not connect: %v", err)
	   	}

	   defer conn.Close()

	   cc := pb.NewChupacabraClient(conn)

	   ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	   defer cancel()

	   rr, err := cc.EnqueueSubmit(ctx, &pb.EnqueueRequest{ExecuteTurn: 111, Message: "woot", Owner: "guy"})

	   	if err != nil {
	   		log.Fatalf("could not greet: %v", err)
	   	}

	   log.Printf("Result: %s", rr.GetToken())
	*/
}
