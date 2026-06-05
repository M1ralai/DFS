package db

import (
	"fmt"

	"github.com/M1ralai/DFS/node/src/utils/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	return sqlx.Connect("postgres", dsn)
}

// Migrate is a no-op until node modules (heartbeat/chunk/ack) introduce their own tables.
func Migrate(db *sqlx.DB, isDebug bool) error {
	_ = db
	_ = isDebug
	return nil
}
