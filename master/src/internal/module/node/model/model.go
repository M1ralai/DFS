package model

import (
	"time"

	"github.com/M1ralai/DFS/src/utils/pg"
	"github.com/google/uuid"
)

type NodeStatus string

const (
	StatusLive     NodeStatus = "live"
	StatusDead     NodeStatus = "dead"
	StatusAssigned NodeStatus = "assigned"
)

type Node struct {
	ID             uuid.UUID    `db:"id"`
	AvailableSpace int          `db:"available_space"`
	Status         NodeStatus   `db:"status"`
	LastHeartbeat  time.Time    `db:"last_heartbeat"`
	Chunks         pg.UUIDArray `db:"chunks"`
}
