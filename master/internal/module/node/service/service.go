package service

import (
	"context"
	"time"

	"github.com/M1ralai/DFS/master/internal/module/node/dto"
	"github.com/M1ralai/DFS/master/internal/module/node/model"
	"github.com/M1ralai/DFS/master/internal/module/node/repository"
	"github.com/M1ralai/DFS/master/utils/config"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
)

type INodeCommService interface {
	Save(dto.NodeSaveRequest) error
	FindAll() ([]model.Node, error)
	HearthBeat(dto.HearthBeatRequest) error
	Acknowledgement(dto.AckRequest) error
	StartDeadNodeChecker(ctx context.Context)
}

type NodeCommService struct {
	master model.Master
	repo   repository.INodeCommRepository
	cfg    config.NodeCommConfig
}

func NewNodeCommService(repo repository.INodeCommRepository, cfg config.NodeCommConfig) INodeCommService {
	return &NodeCommService{
		master: model.Master{
			Nodes:    make(map[uuid.UUID]*model.Node),
			ChunkMap: make(map[uuid.UUID][]uuid.UUID),
			AckCount: make(map[uuid.UUID]int),
		},
		repo: repo,
		cfg:  cfg,
	}
}

func (s *NodeCommService) Save(node dto.NodeSaveRequest) error {
	n := model.Node{
		ID:             node.ID,
		AvailableSpace: node.AvailableSpace,
		Status:         model.StatusLive,
		LastHearthbeat: time.Now(),
		Chunks:         node.Chunks,
	}
	return s.repo.Save(n)
}

func (s *NodeCommService) FindAll() ([]model.Node, error) {
	return s.repo.FindAll()
}

func (s *NodeCommService) HearthBeat(req dto.HearthBeatRequest) error {
	n := model.Node{
		ID:             req.NodeId,
		AvailableSpace: req.AvailableSpace,
	}
	return s.repo.UpdateHearthbeat(n)
}

func (s *NodeCommService) Acknowledgement(req dto.AckRequest) error {
	n := model.Node{
		ID:             req.NodeId,
		AvailableSpace: req.AvailableSpace,
	}
	if err := s.repo.UpdateHearthbeat(n); err != nil {
		return err
	}
	s.master.AckCount[req.ChunkId] += 1
	if s.master.AckCount[req.ChunkId] == s.cfg.ReplicationFactor {
		delete(s.master.AckCount, req.ChunkId)
		log.Infow("chunk fully replicated", "chunkID", req.ChunkId)
	}
	return nil
}

func (s *NodeCommService) StartDeadNodeChecker(ctx context.Context) {
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
