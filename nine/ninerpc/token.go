package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nine/nineora"
)

type TokenQueryReq struct {
	Page *pagination.Page `json:"page"`
}

type TokenQueryByNetworkReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
	Page      *pagination.Page  `json:"page"`
}

type TokenGetReq struct {
	NID nineora.TokenID `json:"nid"`
}

const (
	TokenPath               = "token"
	TokenQueryPath          = TokenPath + "/q"
	TokenQueryByNetworkPath = TokenPath + "/q/by_network"
	TokenGetPath            = TokenPath + "/get"
)

type TokenQuery func(req *TokenQueryReq) (*pagination.Pagination[nineora.Token], *errors.Error)
type TokenQueryByNetwork func(req *TokenQueryByNetworkReq) (*pagination.Pagination[nineora.Token], *errors.Error)
type TokenGet func(req *TokenGetReq) (*nineora.Token, *errors.Error)
