package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("KAFKA_BROKERS", "")
	cfg := Load()
	if cfg.KafkaTopic != "order.completed" || cfg.KafkaGroupID != "wh-suggestions" {
		t.Fatalf("unexpected defaults: %+v", cfg)
	}
	if len(cfg.KafkaBrokers) != 1 || cfg.KafkaBrokers[0] != "redpanda:9092" {
		t.Fatalf("unexpected brokers: %v", cfg.KafkaBrokers)
	}
}

func TestLoadOverrides(t *testing.T) {
	t.Setenv("KAFKA_BROKERS", "a:1,b:2")
	t.Setenv("JWT_SECRET", "s")
	cfg := Load()
	if len(cfg.KafkaBrokers) != 2 || cfg.JWTSecret != "s" {
		t.Fatalf("overrides ignored: %+v", cfg)
	}
}
