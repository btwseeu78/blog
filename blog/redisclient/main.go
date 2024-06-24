package main

import (
	"flag"
	"fmt"
	pb "github.com/btwseeu78/blog/blog/proto"
)

var addr = flag.String("addr", "localhost:6379", "redis server address")

func main() {
	flag.Parse()
	user := &pb.Blog{
		Id:       "2",
		AuthorId: "200",
		Totle:    "test",
		Content:  "test2",
	}
	bcl := GetRedisClient(*addr)
	pong, err := bcl.CheckConnection()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	err = bcl.SetData(user)
	if err != nil {
		fmt.Println(err)
	}
	out := bcl.GetData(user.Id)
	fmt.Println(out)
	out2, err := bcl.ListData(user.Id)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range out2 {
		fmt.Println(val)
	}

}
