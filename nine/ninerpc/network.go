package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nine/nineora"
)

type NetworkQueryReq struct {
	Page *pagination.Page `json:"page"`
}

type NetworkGetReq struct {
	NID nineora.NetworkID `json:"nid"`
}

const (
	NetworkPath      = "network"
	NetworkQueryPath = NetworkPath + "/q"
	NetworkGetPath   = NetworkPath + "/get"
)

type NetworkQuery func(req *NetworkQueryReq) (*pagination.Pagination[nineora.Network], *errors.Error)
type NetworkGet func(req *NetworkGetReq) (*nineora.Network, *errors.Error)
