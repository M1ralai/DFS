package repository

import (
	"time"

	"github.com/M1ralai/DFS/master/internal/module/node/model"
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) IRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(node model.Node) error {
	if _, err := r.db.NamedExec(`INSERT INTO nodes (id, available_space, status, last_hearthbeat, chunks)
	VALUES(:id, :available_space,:status,:last_hearthbeat,:chunks);`, node); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) FindAll() ([]model.Node, error) {
	v := new([]model.Node)
	if err := r.db.Select(v, "SELECT * FROM nodes;"); err != nil {
		return nil, err
	}
	return *v, nil
}

func (r *PostgresRepository) UpdateHearthbeat(n model.Node) error {
	if _, err := r.db.NamedExec(`UPDATE nodes SET last_hearthbeat = NOW(), available_space = :available_space WHERE id = :id`, n); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) FindExpiredNode(t time.Duration) ([]string, error) {
	v := new([]string)
	if err := r.db.Select(v, `SELECT id FROM nodes
WHERE last_hearthbeat < NOW() - $1::interval
AND status != 'dead';`, int(t.Seconds())); err != nil {
		return nil, err
	}
	return *v, nil
}

func (r *PostgresRepository) MarkAsDead(id string) error {
	if _, err := r.db.Exec(`UPDATE nodes SET status = 'dead' WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
