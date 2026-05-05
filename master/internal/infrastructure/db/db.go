package db

import (
	"fmt"

	"github.com/M1ralai/DFS/master/utils/config"
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
		if _, err := db.Exec(`DROP TABLE IF EXISTS nodes`); err != nil {
			return err
		}
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS nodes (
    id VARCHAR PRIMARY KEY,
    available_space INTEGER,
    status VARCHAR,
    last_hearthbeat TIMESTAMP,
    chunks TEXT[]
);
`); err != nil {
		return err
	}
	return nil
}
