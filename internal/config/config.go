// Package config loads the service configuration from the environment.
package config

import (
	"os"
	"strings"
)

// Config holds every setting the service reads at startup.
type Config struct {
	HTTPAddr     string
	RedisAddr    string
	KafkaBrokers []string
	KafkaTopic   string
	KafkaGroupID string
	JWTSecret    string
}

// Load reads the configuration from environment variables, applying defaults
// that match docker-compose.yml.
func Load() Config {
	return Config{
		HTTPAddr:     get("HTTP_ADDR", ":8080"),
		RedisAddr:    get("REDIS_ADDR", "redis:6379"),
		KafkaBrokers: strings.Split(get("KAFKA_BROKERS", "redpanda:9092"), ","),
		KafkaTopic:   get("KAFKA_TOPIC", "order.completed"),
		KafkaGroupID: get("KAFKA_GROUP_ID", "wh-suggestions"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}
}

func get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
