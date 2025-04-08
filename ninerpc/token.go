package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nineora"
)

type TokensReq struct {
	Page pagination.Page `json:"page"`
}

type NetworkTokensReq struct {
	NetworkID nineora.NetworkID `json:"network_id"`
}

type TokenService interface {
	Tokens(req TokensReq) (pagination.Pagination[nineora.Token], *errors.Error)
	NetworkTokens(req NetworkTokensReq) ([]*nineora.Token, *errors.Error)
	Token(id nineora.TokenID) (*nineora.Token, *errors.Error)
}
