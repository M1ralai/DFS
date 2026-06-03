package config

import (
	"time"

	"github.com/M1ralai/DFS/src/utils/env"
)

type NodeCommConfig struct {
	NodeTimeout       time.Duration
	ReplicationFactor int
	ChunkSize         int64
}

func newNodeCommConfig() NodeCommConfig {
	return NodeCommConfig{
		NodeTimeout:       time.Duration(env.IntGetEnv("NODE_TIMEOUT", 5)) * time.Second,
		ReplicationFactor: env.IntGetEnv("REPLICATION_FACTOR", 2),
		ChunkSize:         int64(env.IntGetEnv("CHUNK_SIZE", 32)),
	}
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func newDBConfig() DBConfig {
	return DBConfig{
		Host:     env.GetEnv("DB_HOST", "localhost"),
		Port:     env.GetEnv("DB_PORT", "5432"),
		User:     env.GetEnv("DB_USER", "postgres"),
		Password: env.GetEnv("DB_PASSWORD", "postgres"),
		DBName:   env.GetEnv("DB_NAME", "master"),
		SSLMode:  env.GetEnv("DB_SSL_MODE", "disable"),
	}
}

type Config struct {
	Port              string
	Host              string
	HeartbeatInterval time.Duration
	NodeCommCfg       NodeCommConfig
	DBCfg             DBConfig
}

func LoadConfig() *Config {
	return &Config{
		Port:              env.GetEnv("PORT", ":3030"),
		Host:              env.GetEnv("HOST", "0.0.0.0"),
		HeartbeatInterval: time.Duration(env.IntGetEnv("HEARTHBEAT_INTERVAL", 5)) * time.Second,
		NodeCommCfg:       newNodeCommConfig(),
		DBCfg:             newDBConfig(),
	}
}
