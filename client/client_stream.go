package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vinith/grpc-demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming is started")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names:%v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)

		}
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streamuing is finished")
	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)

}
