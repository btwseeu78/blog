package redisclient

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type BlogClient struct {
	redisClient *redis.Client
}

func GetRedisClient(addr string) *BlogClient {
	options := &redis.Options{
		Addr: addr,
	}
	client := redis.NewClient(options)
	return &BlogClient{
		redisClient: client,
	}
}

func (b *BlogClient) ListData() {
	fmt.Sprintf("LisData")
}
func (b *BlogClient) GetData(key string) string {
	fmt.Sprintf("GetData")
	return "test"
}
func (b *BlogClient) DeleteData(key string) {
	fmt.Sprintf("DeleteData")
}
