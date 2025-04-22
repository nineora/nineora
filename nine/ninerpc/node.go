package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nine/nineora"
)

type NodeQueryByNetworkReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
	Page      *pagination.Page  `json:"page"`
	WithCore  bool              `json:"with_core"`
	Deep      []uint32          `json:"deep"`
}
type NodeQueryBySuperiorReq struct {
	Superior nineora.NodeID   `json:"superior"`
	Page     *pagination.Page `json:"page"`
	WithCore bool             `json:"with_core"`
	Deep     []uint32         `json:"deep"`
}
type NodeGetReq struct {
	NID nineora.NodeID `json:"nid"`
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
