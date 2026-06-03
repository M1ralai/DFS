package db

import (
	"fmt"

	"github.com/M1ralai/DFS/src/utils/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	return sqlx.Connect("postgres", dsn)
}

func Migrate(db *sqlx.DB, isDebug bool) error {
	if isDebug {
		if _, err := db.Exec(`DROP TABLE IF EXISTS chunks`); err != nil {
			return err
		}
		if _, err := db.Exec(`DROP TABLE IF EXISTS files`); err != nil {
			return err
		}
		if _, err := db.Exec(`DROP TABLE IF EXISTS nodes`); err != nil {
			return err
		}
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS nodes (
    id VARCHAR PRIMARY KEY,
    available_space INTEGER,
    status VARCHAR,
    last_heartbeat TIMESTAMP,
    chunks TEXT[]
);
`); err != nil {
		return err
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS files (
    id VARCHAR PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    filename TEXT NOT NULL,
    size BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
`); err != nil {
		return err
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS chunks (
    id VARCHAR PRIMARY KEY,
    file_id VARCHAR NOT NULL REFERENCES files(id) ON DELETE CASCADE,
    chunk_index INTEGER NOT NULL,
    nodes TEXT[] NOT NULL,
    replica_count INTEGER NOT NULL DEFAULT 0,
    UNIQUE(file_id, chunk_index)
);
CREATE INDEX IF NOT EXISTS idx_chunks_file_id ON chunks(file_id);
`); err != nil {
		return err
	}
	return nil
}
