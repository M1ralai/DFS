package repository

import (
	"time"

	"github.com/M1ralai/DFS/src/internal/module/node/model"
	"github.com/google/uuid"
)

type INodeRepository interface {
	Save(model.Node) error
	FindAll() ([]model.Node, error)
	UpdateHeartbeat(model.Node) error
	FindExpiredNode(time.Duration) ([]uuid.UUID, error)
	MarkAsDead(id uuid.UUID) error
}
