package main

import (
"log"
"net"

pb "github.com/alknopfler/testGRPC/requester"
"golang.org/x/net/context"
"google.golang.org/grpc"
	"errors"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Process(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	m := make(map[string]string)
	m["key1"]="value1"
	m["key2"]="value2"

	if (in.KeyId == "key1" || in.KeyId == "key2") {
		return &pb.Response{Keyvalue: m}, nil
	}
	return &pb.Response{Keyvalue: m}, errors.New("not found")

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRequesterServer(s, &server{})
	s.Serve(lis)
}
