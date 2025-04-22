package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nine/nineora"
)

type BusinessQueryReq struct {
	WithCore bool             `json:"with_core"`
	Page     *pagination.Page `json:"page"`
}

type BusinessQueryByNetworkReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
	WithCore  bool              `json:"with_core"`
	Page      *pagination.Page  `json:"page"`
}

type BusinessGetReq struct {
	NID nineora.BusinessID `json:"nid"`
}

const (
	BusinessPath               = "biz"
	BusinessQueryPath          = BusinessPath + "/q"
	BusinessQueryByNetworkPath = BusinessPath + "/q/by_network"
	BusinessGetPath            = BusinessPath + "/get"
)

type BusinessQuery func(req *BusinessQueryReq) (*pagination.Pagination[nineora.Business], *errors.Error)
type BusinessQueryByNetwork func(req *BusinessQueryByNetworkReq) (pagination.Pagination[nineora.Business], *errors.Error)
type BusinessGet func(req *BusinessGetReq) (*nineora.Business, *errors.Error)
