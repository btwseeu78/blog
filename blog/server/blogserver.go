package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/btwseeu78/blog/blog/proto"
	rs "github.com/btwseeu78/blog/blog/redisclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var redisAddr = flag.String("redisAddr", "localhost:6379", "redis server address")

func (s *server) CreateBlog(ctx context.Context, blog *pb.Blog) (newId *pb.BlogId, err error) {
	fmt.Println("CreateBlog Method From Server is called")
	flag.Parse()
	client := rs.GetRedisClient(*redisAddr)
	fmt.Println("CreateBlog Method From Redis Client is called")

	// create the blog
	newId, err = client.SetData(blog)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Printf("CreateBlog New Id is %s\n", newId)
	return
}
