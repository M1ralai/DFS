package model

import (
	"github.com/M1ralai/DFS/src/utils/pg"
	"github.com/google/uuid"
)

type File struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	FileName string    `db:"filename"`
	FileSize int64     `db:"size"`
}

type Chunk struct {
	ID           uuid.UUID    `db:"id"`
	FileID       uuid.UUID    `db:"file_id"`
	ChunkIndex   int          `db:"chunk_index"`
	Nodes        pg.UUIDArray `db:"nodes"`
	ReplicaCount int          `db:"replica_count"`
}
