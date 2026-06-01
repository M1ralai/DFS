package dto

import "github.com/google/uuid"

type HearthBeatRequest struct {
	NodeId         uuid.UUID `json:"node_id" validate:"required"`
	AvailableSpace int       `json:"available_space" validate:"required"`
}

type AckRequest struct {
	NodeId         uuid.UUID `json:"id" validate:"required"`
	ChunkId        uuid.UUID `json:"chunk_id" validate:"required"`
	AvailableSpace int       `json:"available_space" validate:"required"`
}

type NodeSaveRequest struct {
	ID             uuid.UUID   `json:"node_id" validate:"required"`
	AvailableSpace int         `json:"available_space" validate:"required"`
	Chunks         []uuid.UUID `json:"chunks" validate:"required,omitnil"`
}
