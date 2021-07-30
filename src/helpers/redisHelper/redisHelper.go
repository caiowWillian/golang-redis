package redisHelper

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Set() {
	ctx := context.Background()
	rdb = getClient()
	rdb.Set(ctx, "key", "haha", 0)
	val := rdb.Get(ctx, "key").Val()
	fmt.Println(val)
}

func getClient() *redis.Client {
	if rdb != nil {
		fmt.Println("teste")
		return rdb
	}

	return redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
}
