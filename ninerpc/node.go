package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nineora"
)

type NodeQueryByNetworkReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
	Page      *pagination.Page  `json:"page"`
}
type NodeQueryBySuperiorReq struct {
	Superior nineora.NodeID   `json:"superior"`
	Page     *pagination.Page `json:"page"`
}
type NodeGetReq struct {
	ID nineora.NodeID `json:"id"`
}

const (
	NodePath                = "node"
	NodeQueryByNetworkPath  = NodePath + "/q/by_network"
	NodeQueryBySuperiorPath = NodePath + "/q/by_superior"
	NodeGetPath             = NodePath + "/get"
)

type NodeQueryByNetwork func(req *NodeQueryByNetworkReq) (*pagination.Pagination[nineora.Node], *errors.Error)
type NodeQueryBySuperior func(req *NodeQueryBySuperiorReq) (*pagination.Pagination[nineora.Node], *errors.Error)
type NodeGet func(req *NodeGetReq) (*nineora.Node, *errors.Error)
