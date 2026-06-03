package dto

import "github.com/google/uuid"

type HeartbeatRequest struct {
	ID             uuid.UUID `json:"node_id" validate:"required"`
	AvailableSpace int       `json:"available_space" validate:"required"`
}

type AckRequest struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	ChunkID        uuid.UUID `json:"chunk_id" validate:"required"`
	AvailableSpace int       `json:"available_space" validate:"required"`
}

type NodeSaveRequest struct {
	ID             uuid.UUID   `json:"node_id" validate:"required"`
	AvailableSpace int         `json:"available_space" validate:"required"`
	Chunks         []uuid.UUID `json:"chunks" validate:"required,omitnil"`
}
