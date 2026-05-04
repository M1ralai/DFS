package handler

import "github.com/M1ralai/DFS/master/internal/module/node/service"

type NodeCommHandler struct {
	service service.INodeCommService
}

func NewNodeCommHandler() NodeCommHandler {
	return NodeCommHandler{
		service: service.NewNodeCommService(),
	}
}
