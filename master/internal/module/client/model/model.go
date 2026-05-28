package model

import "github.com/google/uuid"

type File struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	FileName string    `db:"filename"`
	FileSize int64     `db:"size"`
}

type Chunk struct {
	ID     uuid.UUID `db:"id"`
	FileID string    `db:"file_id"`
	Nodes  []string  `db:"nodes"`
}
