package model

import (
	"time"

	"github.com/lib/pq"
)

type NodeStatus string

const (
	StatusLive     NodeStatus = "live"
	StatusDead     NodeStatus = "dead"
	StatusAssigned NodeStatus = "assigned"
)

type Node struct {
	ID             string         `db:"id"`
	AvailableSpace int            `db:"available_space"`
	Status         NodeStatus     `db:"status"`
	LastHearthbeat time.Time      `db:"last_hearthbeat"`
	Chunks         pq.StringArray `db:"chunks"`
}

type Master struct {
	Nodes    map[string]*Node
	ChunkMap map[string][]string
	AckCount map[string]int
}
