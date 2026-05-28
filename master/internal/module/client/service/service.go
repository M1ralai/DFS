package service

import "github.com/M1ralai/DFS/master/internal/module/client/repository"

type IClientService interface {
}

type ClientService struct {
	repo *repository.ClientCommRepository
}

func NewClientService(repo *repository.ClientCommRepository)
