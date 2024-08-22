package db

import (
    "context"
    "github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init (addr string) {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
}

func WriteUrl (url string, key string) (string, error) {
	ctx := context.Background()
	return key, Client.Set(ctx, key, url, 0).Err()
}

func ReadUrl (key string) (string, error) {
	ctx := context.Background()
	return Client.Get(ctx, key).Result()
}