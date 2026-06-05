package config

import (
	"time"

	"github.com/M1ralai/DFS/node/src/utils/env"
)

type NodeConfig struct {
	NodeTimeout       time.Duration
	ReplicationFactor int
	ChunkSize         int64
}

func newNodeConfig() NodeConfig {
	return NodeConfig{
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
		DBName:   env.GetEnv("DB_NAME", "node"),
		SSLMode:  env.GetEnv("DB_SSL_MODE", "disable"),
	}
}

type Config struct {
	Port              string
	Host              string
	HeartbeatInterval time.Duration
	NodeCfg           NodeConfig
	DBCfg             DBConfig
	StorageDir        string
	MasterURL         string
}

func LoadConfig() *Config {
	return &Config{
		Port:              env.GetEnv("PORT", ":4040"),
		Host:              env.GetEnv("HOST", "0.0.0.0"),
		HeartbeatInterval: time.Duration(env.IntGetEnv("HEARTBEAT_INTERVAL", 5)) * time.Second,
		NodeCfg:           newNodeConfig(),
		DBCfg:             newDBConfig(),
		StorageDir:        env.GetEnv("STORAGE_DIR", "./data/chunks"),
		MasterURL:         env.GetEnv("MASTER_URL", "http://localhost:3030"),
	}
}
