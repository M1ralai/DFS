package service

import (
	"fmt"
	"math"

	"github.com/M1ralai/DFS/src/internal/module/client/dto"
	"github.com/M1ralai/DFS/src/internal/module/client/model"
	ClientRepo "github.com/M1ralai/DFS/src/internal/module/client/repository"
	nodeModel "github.com/M1ralai/DFS/src/internal/module/node/model"
	NodeRepo "github.com/M1ralai/DFS/src/internal/module/node/repository"
	"github.com/M1ralai/DFS/src/utils/config"
	"github.com/google/uuid"
)

type IClientService interface {
	UploadFile(dto.UploadRequest) (dto.UploadResponse, error)
	GetFile(id uuid.UUID) (dto.FileResponse, error)
	DeleteFile(req dto.DeleteRequest) error
	GetUserFiles(userID uuid.UUID) (dto.UserFilesResponse, error)
}

type ClientService struct {
	repo     ClientRepo.IClientRepository
	nodeRepo NodeRepo.INodeRepository
	cfg      config.NodeConfig
}

func NewClientService(clientRepo ClientRepo.IClientRepository, nodeRepo NodeRepo.INodeRepository, cfg config.NodeConfig) IClientService {
	return &ClientService{
		repo:     clientRepo,
		nodeRepo: nodeRepo,
		cfg:      cfg,
	}
}

func (s *ClientService) UploadFile(req dto.UploadRequest) (dto.UploadResponse, error) {
	fID, err := uuid.NewUUID()
	if err != nil {
		return dto.UploadResponse{}, err
	}

	chunkCount := int(math.Ceil(float64(req.FileSize) / float64(s.cfg.ChunkSize)))

	allNodes, err := s.nodeRepo.FindAll()
	if err != nil {
		return dto.UploadResponse{}, err
	}

	liveNodes := make([]nodeModel.Node, 0)
	for _, n := range allNodes {
		if n.Status == nodeModel.StatusLive {
			liveNodes = append(liveNodes, n)
		}
	}

	if len(liveNodes) < s.cfg.ReplicationFactor {
		return dto.UploadResponse{}, fmt.Errorf(
			"yetersiz canlı node: %d live, %d gerekli (replication factor)",
			len(liveNodes), s.cfg.ReplicationFactor,
		)
	}

	f := model.File{
		ID:       fID,
		UserID:   req.UserID,
		FileName: req.FileName,
		FileSize: req.FileSize,
	}

	if err := s.repo.PostFile(f); err != nil {
		return dto.UploadResponse{}, err
	}

	rf := s.cfg.ReplicationFactor
	chunks := make([]dto.ChunkLocation, 0, chunkCount)
	nodeIdx := 0

	for i := range chunkCount {
		cID, err := uuid.NewUUID()
		if err != nil {
			return dto.UploadResponse{}, err
		}

		selectedNodes := make([]uuid.UUID, 0, rf)
		for range rf {
			selectedNodes = append(selectedNodes, liveNodes[nodeIdx%len(liveNodes)].ID)
			nodeIdx++
		}

		c := model.Chunk{
			ID:         cID,
			FileID:     fID,
			ChunkIndex: i,
			Nodes:      selectedNodes,
		}

		if err := s.repo.PostChunk(c); err != nil {
			return dto.UploadResponse{}, err
		}

		chunks = append(chunks, dto.ChunkLocation{
			ChunkID: cID,
			Nodes:   selectedNodes,
		})
	}

	return dto.UploadResponse{
		FileID:  fID,
		Chunks:  chunks,
		Message: "upload organize edildi, client chunk'lari ilgili node'lara gondermeli",
	}, nil
}

func (s *ClientService) GetFile(id uuid.UUID) (dto.FileResponse, error) {
	f, err := s.repo.GetFile(id)
	if err != nil {
		return dto.FileResponse{}, err
	}

	chunks, err := s.repo.GetChunksByFileID(f.ID)
	if err != nil {
		return dto.FileResponse{}, err
	}

	chunkLocations := make([]dto.ChunkLocation, 0, len(chunks))
	for _, c := range chunks {
		chunkLocations = append(chunkLocations, dto.ChunkLocation{
			ChunkID: c.ID,
			Nodes:   c.Nodes,
		})
	}

	return dto.FileResponse{
		FileID:   f.ID,
		FileName: f.FileName,
		Size:     f.FileSize,
		UserID:   f.UserID,
		Chunks:   chunkLocations,
	}, nil
}

func (s *ClientService) DeleteFile(req dto.DeleteRequest) error {
	if err := s.repo.DeleteChunksByFileID(req.FileID); err != nil {
		return err
	}
	if err := s.repo.DeleteFile(req.FileID); err != nil {
		return err
	}
	return nil
}

// GetUserFiles — kullanıcının tüm dosyalarını listeler.
func (s *ClientService) GetUserFiles(userID uuid.UUID) (dto.UserFilesResponse, error) {
	files, err := s.repo.GetAllFileUser(userID)
	if err != nil {
		return dto.UserFilesResponse{}, err
	}

	responses := make([]dto.FileResponse, 0, len(files))
	for _, f := range files {
		chunks, err := s.repo.GetChunksByFileID(f.ID)
		if err != nil {
			return dto.UserFilesResponse{}, err
		}

		chunkLocations := make([]dto.ChunkLocation, 0, len(chunks))
		for _, c := range chunks {
			chunkLocations = append(chunkLocations, dto.ChunkLocation{
				ChunkID: c.ID,
				Nodes:   c.Nodes,
			})
		}

		responses = append(responses, dto.FileResponse{
			FileID:   f.ID,
			FileName: f.FileName,
			Size:     f.FileSize,
			UserID:   f.UserID,
			Chunks:   chunkLocations,
		})
	}

	return dto.UserFilesResponse{Files: responses}, nil
}
