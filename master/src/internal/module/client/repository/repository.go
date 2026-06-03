package repository

import (
	"github.com/M1ralai/DFS/src/internal/module/client/model"
	"github.com/google/uuid"
)

type IClientRepository interface {
	PostFile(f model.File) error
	GetFile(id uuid.UUID) (model.File, error)
	GetAllFileUser(id uuid.UUID) ([]model.File, error)
	DeleteFile(id uuid.UUID) error

	PostChunk(c model.Chunk) error
	GetChunk(id uuid.UUID) (model.Chunk, error)
	GetChunksByFileID(id uuid.UUID) ([]model.Chunk, error)
	DeleteChunksByFileID(id uuid.UUID) error
	DeleteChunk(id uuid.UUID) error

	IncrementReplicaCount(chunkID uuid.UUID) (int, error)
}
