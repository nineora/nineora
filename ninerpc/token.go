package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nineora"
)

type TokenQueryReq struct {
	Page *pagination.Page `json:"page"`
}

type TokenQueryByNetworkReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
}

type TokenGetReq struct {
	ID nineora.TokenID `json:"id"`
}

const (
	TokenPath               = "token"
	TokenQueryPath          = TokenPath + "/q"
	TokenQueryByNetworkPath = TokenPath + "/q/by_network"
	TokenGetPath            = TokenPath + "/get"
)

type TokenQuery func(req *TokenQueryReq) (*pagination.Pagination[nineora.Token], *errors.Error)
type TokenQueryByNetwork func(req *TokenQueryByNetworkReq) ([]*nineora.Token, *errors.Error)
type TokenGet func(req *TokenGetReq) (*nineora.Token, *errors.Error)
