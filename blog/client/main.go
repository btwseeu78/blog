package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/btwseeu78/blog/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var addr = flag.String("addr", "localhost:50051", "http service address")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("could not close connection: %v", err)
		}
	}(conn)

	c := pb.NewBlogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	CreateBlog(ctx, c)
}
