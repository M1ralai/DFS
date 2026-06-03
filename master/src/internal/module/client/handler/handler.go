package handler

import "github.com/M1ralai/DFS/src/internal/module/client/service"

type ClientHandler struct {
	service service.IClientService
}

func NewClientHandler(service service.IClientService) ClientHandler {
	return ClientHandler{
		service: service,
	}
}
