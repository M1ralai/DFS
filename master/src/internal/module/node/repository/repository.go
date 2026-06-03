package repository

import (
	"time"

	"github.com/M1ralai/DFS/src/internal/module/node/model"
	"github.com/google/uuid"
)

type INodeCommRepository interface {
	Save(model.Node) error
	FindAll() ([]model.Node, error)
	UpdateHearthbeat(model.Node) error
	FindExpiredNode(time.Duration) ([]uuid.UUID, error)
	MarkAsDead(id uuid.UUID) error
}
