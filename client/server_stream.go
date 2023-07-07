package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vinith/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names")
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error while streaming %v", err)
		}
		log.Println(message)

	}
	log.Printf("streaming finished")
}
