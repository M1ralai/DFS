package dto

import "github.com/google/uuid"

type AckRequest struct {
	ID             uuid.UUID
	ChunkID        uuid.UUID
	AvailableSpace int
}

type HeartBeatRequest struct {
	ID             uuid.UUID
	AvailableSpace int
}

type RegisterRequest struct {
	Chuks          []uuid.UUID `json:"chunks"`
	AvailableSpace int         `json:"avilable_space"`
	ID             uuid.UUID   `json:"node_id"`
}
