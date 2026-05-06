package repository

import "github.com/M1ralai/DFS/master/internal/module/client/model"

type IClientCommRepository interface {
	PostFile(f model.File) error
	GetFile(id string) (model.File, error)
	GetAllFileUser(id string) ([]model.File, error)
	DeleteFile(id string) error

	PostChunk(c model.Chunk) error
	GetChunk(id string) (model.Chunk, error)
	DeleteChunk(id string) error
}
