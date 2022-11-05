package main

import (
	"context"
	"github.com/cyj19/gowalk"
	"github.com/cyj19/gowalk/component/redis"
	"log"
	"time"
)

func main() {
	gowalk.Run(&redis.Instance{})
	key := "gowalk"
	v := "cyj19"
	ctx := context.Background()
	_ = redis.Main().Set(ctx, key, v, 10*time.Second)
	result, _ := redis.Main().Get(ctx, key).Result()
	log.Println(result)
}
