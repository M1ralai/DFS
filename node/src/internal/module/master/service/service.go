package service

import (
	"github.com/M1ralai/DFS/node/src/internal/module/master/client"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type IMasterService interface {
	Register() error
	StartHeartBeat(fiber.Ctx)
	AcknowledgeChunk(uuid.UUID) error
	AvailableSpace() int
}

type MasterService struct {
	masterClient client.MasterClient
	/**
	*	IChunkRepo
	*	SpaceProvider
	*	Config
	* **/
}

func NewMasterService(masterClient client.MasterClient) IMasterService {
	return &MasterService{
		masterClient: masterClient,
	}
}

func (s *MasterService) Register() error {
	return nil
}

func (s *MasterService) StartHeartBeat(ctx fiber.Ctx) {}

func (s *MasterService) AcknowledgeChunk(ChunkID uuid.UUID) error {
	return nil
}

func (s *MasterService) AvailableSpace() int {
	return 0
}
