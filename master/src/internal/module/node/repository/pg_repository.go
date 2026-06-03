package repository

import (
	"time"

	"github.com/M1ralai/DFS/src/internal/module/node/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type NodeRepository struct {
	db *sqlx.DB
}

func NewNodeRepository(db *sqlx.DB) INodeRepository {
	return &NodeRepository{
		db: db,
	}
}

func (r *NodeRepository) Save(node model.Node) error {
	if _, err := r.db.NamedExec(`INSERT INTO nodes (id, available_space, status, last_heartbeat, chunks)
	VALUES(:id, :available_space,:status,:last_heartbeat,:chunks);`, node); err != nil {
		return err
	}
	return nil
}

func (r *NodeRepository) FindAll() ([]model.Node, error) {
	v := new([]model.Node)
	if err := r.db.Select(v, "SELECT * FROM nodes;"); err != nil {
		return nil, err
	}
	return *v, nil
}

func (r *NodeRepository) UpdateHeartbeat(n model.Node) error {
	if _, err := r.db.NamedExec(`UPDATE nodes SET last_heartbeat = NOW(), available_space = :available_space WHERE id = :id`, n); err != nil {
		return err
	}
	return nil
}

func (r *NodeRepository) FindExpiredNode(t time.Duration) ([]uuid.UUID, error) {
	v := new([]uuid.UUID)
	if err := r.db.Select(v, `SELECT id FROM nodes
WHERE last_heartbeat < NOW() - $1::interval
AND status != 'dead';`, int(t.Seconds())); err != nil {
		return nil, err
	}
	return *v, nil
}

func (r *NodeRepository) MarkAsDead(id uuid.UUID) error {
	if _, err := r.db.Exec(`UPDATE nodes SET status = 'dead' WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
