package configs

import (
	"fmt"
	"os"
)

type Config struct {
	RedisAddress  string
	RedisPassword string
}

func LoadConfig() Config {
	// return Config{
	// 	RedisAddress:  os.Getenv("REDIS_ADDRESS"),
	// 	RedisPassword: os.Getenv("REDIS_PASSWORD"),
	// }
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		// Default to the Redis service name in Docker Compose
		redisAddress = "redis:6379" // This uses the Docker service name
	}

	// Optional: Get Redis password if needed
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Logging the configuration (optional)
	fmt.Printf("Using Redis address: %s\n", redisAddress)

	return Config{
		RedisAddress:  redisAddress,
		RedisPassword: redisPassword,
	}
}
