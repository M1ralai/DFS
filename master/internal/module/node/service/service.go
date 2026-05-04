package service

type INodeCommService interface {
}

type NodeCommService struct {
}

func NewNodeCommService() INodeCommService {
	return NodeCommService{}
}
