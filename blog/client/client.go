package main

import (
	"context"
	"fmt"
	pb "github.com/btwseeu78/blog/blog/proto"
	"log"
)

func CreateBlog(ctx context.Context, c pb.BlogServiceClient) {

	fmt.Println("CreateBlog")
	blog := &pb.Blog{
		AuthorId: "1",
		Totle:    "The Test",
		Content:  "The Test Set",
	}
	id, err := c.CreateBlog(ctx, blog)
	if err != nil {
		log.Fatalf("CreateBlog: %v", err)
	}
	log.Printf("CreateBlog: %v", id)
}
