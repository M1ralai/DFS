package config

import (
	"time"

	"github.com/M1ralai/DFS/master/utils/env"
)

type Config struct {
	Port              string
	Host              string
	NodeTimeout       time.Duration
	HeartbeatInterval time.Duration
	ChunkSize         int // byte
	ReplicationFactor int
}

func LoadConfig() *Config {
	return &Config{
		Port:              env.GetEnv("PORT", ":3030"),
		Host:              env.GetEnv("HOST", "0.0.0.0"),
		HeartbeatInterval: time.Duration(env.IntGetEnv("HEARTHBEAT_INTERVAL", 5)) * time.Second,
		NodeTimeout:       time.Duration(env.IntGetEnv("NODE_TIMEOUT", 5)) * time.Second,
		ChunkSize:         env.IntGetEnv("CHUNK_SIZE", 32),
		ReplicationFactor: env.IntGetEnv("REPLICATION_FACTOR", 2),
	}
}
