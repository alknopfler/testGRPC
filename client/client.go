package main

import (
	"log"

	pb "github.com/alknopfler/testGRPC/requester"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "key1"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRequesterClient(conn)


	key := defaultName
	r, err := c.Process(context.Background(), &pb.Request{KeyId: key})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Output from server: %s", r.Keyvalue[key])
}
