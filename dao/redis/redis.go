package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

func Init(pass, host string, port, db int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pass,
		DB:       db,
	})
	return client.Ping(context.Background()).Err()
}
