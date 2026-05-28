package dto

import "github.com/google/uuid"

type UploadRequest struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	FileName string    `json:"filename" validate:"required"`
	FileSize int64     `json:"file_size" validate:"required"`
}

type ChunkLocation struct {
	ChunkID string   `json:"chunk_id"`
	Nodes   []string `json:"nodes"`
}

type UploadResponse struct {
	FileID  string          `json:"file_id"`
	Chunks  []ChunkLocation `json:"chunks"`
	Message string          `json:"message"`
}

type FileResponse struct {
	FileID   string          `json:"file_id"`
	FileName string          `json:"filename"`
	Size     int64           `json:"size"`
	UserID   uuid.UUID       `json:"user_id"`
	Chunks   []ChunkLocation `json:"chunks"`
}

type DeleteRequest struct {
	FileID string    `json:"file_id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type UserFilesResponse struct {
	Files []FileResponse `json:"files"`
}
