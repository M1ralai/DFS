package repository

import (
	"time"

	"github.com/M1ralai/DFS/master/internal/module/node/model"
)

type IRepository interface {
	Save(model.Node) error
	FindAll() ([]model.Node, error)
	UpdateHearthbeat(model.Node) error
	FindExpiredNode(time.Duration) ([]string, error)
	MarkAsDead(NodeID string) error
}
