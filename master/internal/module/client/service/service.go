package service

import (
	"fmt"
	"math"

	"github.com/M1ralai/DFS/master/internal/module/client/dto"
	"github.com/M1ralai/DFS/master/internal/module/client/model"
	ClientRepo "github.com/M1ralai/DFS/master/internal/module/client/repository"
	nodeModel "github.com/M1ralai/DFS/master/internal/module/node/model"
	NodeRepo "github.com/M1ralai/DFS/master/internal/module/node/repository"
	"github.com/M1ralai/DFS/master/utils/config"
	"github.com/google/uuid"
)

type IClientService interface {
	UploadFile(dto.UploadRequest) (dto.UploadResponse, error)
	GetFile(id uuid.UUID) (dto.FileResponse, error)
	DeleteFile(req dto.DeleteRequest) error
	GetUserFiles(userID uuid.UUID) (dto.UserFilesResponse, error)
}

type ClientService struct {
	repo     ClientRepo.IClientCommRepository
	nodeRepo NodeRepo.INodeCommRepository
	cfg      config.NodeCommConfig
}

func NewClientService(
	clientRepo ClientRepo.IClientCommRepository,
	nodeRepo NodeRepo.INodeCommRepository,
	cfg config.NodeCommConfig,
) *ClientService {
	return &ClientService{
		repo:     clientRepo,
		nodeRepo: nodeRepo,
		cfg:      cfg,
	}
}

func (s ClientService) UploadFile(req dto.UploadRequest) (dto.UploadResponse, error) {
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

	for range chunkCount {
		cID, err := uuid.NewUUID()
		if err != nil {
			return dto.UploadResponse{}, err
		}

		selectedNodes := make([]string, 0, rf)
		for range rf {
			selectedNodes = append(selectedNodes, liveNodes[nodeIdx%len(liveNodes)].ID)
			nodeIdx++
		}

		c := model.Chunk{
			ID:     cID,
			FileID: fID.String(),
			Nodes:  selectedNodes,
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

func (s ClientService) GetFile(id uuid.UUID) (dto.FileResponse, error) {
	f, err := s.repo.GetFile(id)
	if err != nil {
		return dto.FileResponse{}, err
	}

	// TODO: repository'de GetChunksByFileID yok,
	_ = f // f.FileID kullanılarak chunk'lar çekilecek

	return dto.FileResponse{
		FileID:   f.ID,
		FileName: f.FileName,
		Size:     f.FileSize,
		UserID:   f.UserID,
		Chunks:   nil, // TODO: GetChunksByFileID eklendikten sonra doldur
	}, nil
}

// TODO: önce chunk'ları silmesi gerekiyor — repository GetChunksByFileID + DeleteChunkByFileID eklenmeli.
func (s ClientService) DeleteFile(req dto.DeleteRequest) error {
	if err := s.repo.DeleteFile(req.FileID); err != nil {
		return err
	}
	// TODO: ilgili chunk'ları da sil — repository tamamlanınca eklenir
	return nil
}

// GetUserFiles — kullanıcının tüm dosyalarını listeler.
func (s ClientService) GetUserFiles(userID uuid.UUID) (dto.UserFilesResponse, error) {
	files, err := s.repo.GetAllFileUser(userID)
	if err != nil {
		return dto.UserFilesResponse{}, err
	}

	responses := make([]dto.FileResponse, 0, len(files))
	for _, f := range files {
		responses = append(responses, dto.FileResponse{
			FileID:   f.ID,
			FileName: f.FileName,
			Size:     f.FileSize,
			UserID:   f.UserID,
			Chunks:   nil, // TODO: GetChunksByFileID eklendikten sonra doldur
		})
	}

	return dto.UserFilesResponse{Files: responses}, nil
}
