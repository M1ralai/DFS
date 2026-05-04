package repository

type IRepository interface {
	RegisterNode(NodeId int, NodeIp string)
}
