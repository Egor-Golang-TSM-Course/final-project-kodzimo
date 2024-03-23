package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace with your Redis server address
		Password: "",               // replace with your Redis server password if any
		DB:       0,                // default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
