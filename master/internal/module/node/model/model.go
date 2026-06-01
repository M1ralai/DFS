package model

import (
	"time"

	"github.com/M1ralai/DFS/master/utils/pg"
	"github.com/google/uuid"
)

type NodeStatus string

const (
	StatusLive     NodeStatus = "live"
	StatusDead     NodeStatus = "dead"
	StatusAssigned NodeStatus = "assigned"
)

type Node struct {
	ID             uuid.UUID   `db:"id"`
	AvailableSpace int         `db:"available_space"`
	Status         NodeStatus  `db:"status"`
	LastHearthbeat time.Time   `db:"last_hearthbeat"`
	Chunks         pg.UUIDArray `db:"chunks"`
}

type Master struct {
	Nodes    map[uuid.UUID]*Node
	ChunkMap map[uuid.UUID][]uuid.UUID
	AckCount map[uuid.UUID]int
}
