package dto

type HearthBeatRequest struct {
	NodeId         string `json:"node_id" validate:"required"`
	AvailableSpace int    `json:"available_space" validate:"required"`
}

type AckRequest struct {
	NodeId         string `json:"id" validate:"required"`
	ChunkId        string `json:"chunk_id" validate:"required"`
	AvailableSpace int    `json:"available_space" validate:"required"`
}

type NodeSaveRequest struct {
	ID             string   `json:"node_id" validate:"required"`
	AvailableSpace int      `json:"available_space" validate:"required"`
	Chunks         []string `json:"chunks" validate:"required,omitnil"`
}
