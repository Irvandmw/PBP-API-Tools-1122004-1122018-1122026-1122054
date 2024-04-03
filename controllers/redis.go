package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	ping, err2 := client.Ping(context.Background()).Result()
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println(ping)

	err := saveToken(client, "user1", "123")
	if err != nil {
		log.Fatalf("Error saving token: %v", err)
	}
	fmt.Println("Token saved successfully.")

	// Example: retrieving a token
	token, err := getToken(client, "user1")
	if err != nil {
		log.Fatalf("Error retrieving token: %v", err)
	}
	fmt.Println("Token retrieved:", token)

	// Example: sending a report to an external API (pseudo code)
	// sendReportToWhatsApp(token, reportData)
}

// Function to save token to Redis
func saveToken(client *redis.Client, userID, token string) error {
	ctx := context.Background()
	err := client.Set(ctx, userID+":token", token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Function to retrieve token from Redis
func getToken(client *redis.Client, userID string) (string, error) {
	ctx := context.Background()
	token, err := client.Get(ctx, userID+":token").Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
