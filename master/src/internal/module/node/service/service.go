package service

import (
	"context"
	"time"

	"github.com/M1ralai/DFS/src/internal/module/node/dto"
	"github.com/M1ralai/DFS/src/internal/module/node/model"
	"github.com/M1ralai/DFS/src/internal/module/node/repository"
	ClientRepo "github.com/M1ralai/DFS/src/internal/module/client/repository"
	"github.com/M1ralai/DFS/src/utils/config"
	"github.com/gofiber/fiber/v3/log"
)

type INodeService interface {
	Save(dto.NodeSaveRequest) error
	FindAll() ([]model.Node, error)
	Heartbeat(dto.HeartbeatRequest) error
	Acknowledgement(dto.AckRequest) error
	StartDeadNodeChecker(ctx context.Context)
}

type NodeService struct {
	repo       repository.INodeRepository
	clientRepo ClientRepo.IClientRepository
	cfg        config.NodeConfig
}

func NewNodeService(repo repository.INodeRepository, clientRepo ClientRepo.IClientRepository, cfg config.NodeConfig) INodeService {
	return &NodeService{
		repo:       repo,
		clientRepo: clientRepo,
		cfg:        cfg,
	}
}

func (s *NodeService) Save(node dto.NodeSaveRequest) error {
	n := model.Node{
		ID:             node.ID,
		AvailableSpace: node.AvailableSpace,
		Status:         model.StatusLive,
		LastHeartbeat:  time.Now(),
		Chunks:         node.Chunks,
	}
	return s.repo.Save(n)
}

func (s *NodeService) FindAll() ([]model.Node, error) {
	return s.repo.FindAll()
}

func (s *NodeService) Heartbeat(req dto.HeartbeatRequest) error {
	n := model.Node{
		ID:             req.ID,
		AvailableSpace: req.AvailableSpace,
	}
	return s.repo.UpdateHeartbeat(n)
}

func (s *NodeService) Acknowledgement(req dto.AckRequest) error {
	n := model.Node{
		ID:             req.ID,
		AvailableSpace: req.AvailableSpace,
	}
	if err := s.repo.UpdateHeartbeat(n); err != nil {
		return err
	}
	count, err := s.clientRepo.IncrementReplicaCount(req.ChunkID)
	if err != nil {
		return err
	}
	if count >= s.cfg.ReplicationFactor {
		log.Infow("chunk fully replicated", "chunkID", req.ChunkID, "replica_count", count)
	}
	return nil
}

func (s *NodeService) StartDeadNodeChecker(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(s.cfg.NodeTimeout)
		for {
			select {
			case <-ticker.C:
				e, err := s.repo.FindExpiredNode(s.cfg.NodeTimeout)
				if err != nil {
					log.Warn("dead nodelar kontrol edilirken hata yakalandı ", err)
				}
				for _, v := range e {
					if err := s.repo.MarkAsDead(v); err != nil {
						log.Warn("node status dead yaparken hata yakalandı", err)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}
