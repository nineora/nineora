package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nineora"
)

type AccountQueryReq struct {
	Page        *pagination.Page `json:"page"`
	WithBalance bool             `json:"with_balance"`
}

type AccountGetReq struct {
	ID          nineora.AccountID `json:"id"`
	WithBalance bool              `json:"with_balance"`
}

type BillQueryReq struct {
	Page      *pagination.Page  `json:"page"`
	AccountID nineora.AccountID `json:"account_id"`
	Token     string            `json:"token"`
}

const (
	AccountPath      = "account"
	AccountQueryPath = AccountPath + "/q"
	AccountGetPath   = AccountPath + "/get"
	BillQueryPath    = AccountPath + "/bill/q"
)

type AccountQuery func(req *AccountQueryReq) (*pagination.Pagination[nineora.Account], *errors.Error)
type AccountGet func(req *AccountGetReq) (*nineora.Account, *errors.Error)
type BillQuery func(req *BillQueryReq) (*pagination.Pagination[nineora.Bill], *errors.Error)
