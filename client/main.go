package main

import (
	"log"

	pb "github.com/vinith/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"akhil", "vinith", "sai"},
	}

	//callSayHello(client)

	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectional(client, names)
}
