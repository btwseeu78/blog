package main

import (
	"flag"
	pb "github.com/btwseeu78/blog/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var addr = flag.String("addr", ":8080", "http service address")

type server struct {
	pb.BlogServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	log.Printf("listening on %s", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
