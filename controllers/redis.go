package controllers

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Redis() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0, // use default DB
		})
	}

	return client
}

// Function to save token to Redis
func SaveToken(client *redis.Client, key string, value interface{}) error {
	ctx := context.Background()
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, p, 0).Err()
}

// Function to retrieve token from Redis
func GetToken(client *redis.Client, key string, dest interface{}) error {
	ctx := context.Background()
	p, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}