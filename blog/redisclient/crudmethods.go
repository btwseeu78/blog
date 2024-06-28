package redisclient

import (
	"context"
	"fmt"
	pb "github.com/btwseeu78/blog/blog/proto"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"time"
)

type BlogClient struct {
	redisClient *redis.Client
}

func GetRedisClient(addr string) *BlogClient {
	options := &redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(options)
	return &BlogClient{
		redisClient: client,
	}
}

func (b *BlogClient) CheckConnection() (pong bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//test connection
	success, err := b.redisClient.Ping(ctx).Result()
	if err != nil {
		return false, err
	}
	fmt.Println("success:", success)
	return success == "PONG", nil
}

func (b *BlogClient) GetData(key int64) *pb.Blog {
	searchKey := fmt.Sprintf("blog:%v", key)
	val, err := b.redisClient.Get(context.Background(), searchKey).Result()
	if err != nil {
		fmt.Println("redis unable to get value, err:", err)
	}
	var retrivedUser pb.Blog
	byteval := []byte(val)
	err = proto.Unmarshal(byteval, &retrivedUser)
	if err != nil {
		fmt.Println("redis unable to unmarshal, err:", err)
	}
	return &retrivedUser
}

func (b *BlogClient) DeleteData(key string) error {
	searchKey := fmt.Sprintf("blog:%v", key)
	err := b.redisClient.Del(context.Background(), searchKey).Err()
	if err != nil {
		fmt.Println("redis unable to delete value, err:", err)
		return err
	}
	return nil
}
func (b *BlogClient) SetData(blog *pb.Blog) (out *pb.BlogId, err error) {
	fmt.Println("NewBlogCalled")
	newId, err := b.redisClient.Incr(context.Background(), "blog").Result()
	fmt.Println("newId:", newId)
	if err != nil {
		fmt.Println("redis unable to set value, err:", err)
	}
	data, err := proto.Marshal(blog)
	if err != nil {
		fmt.Printf("Marshal err: %v\n", err)
	}
	key := fmt.Sprintf("blog:%v", newId)
	err = b.redisClient.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		return
	}
	return &pb.BlogId{
		Id: key,
	}, nil

}
