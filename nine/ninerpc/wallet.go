package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/nineora/nineora/nine/ninekey"
	"github.com/nineora/nineora/nine/nineora"
)

type WalletBindReq struct {
	Link nineora.Link           `json:"link"`
	Seed nineora.Seed           `json:"seed"`
	Meta map[string]interface{} `json:"meta"`
}

type AccCreateReq struct {
	Link       nineora.Link           `json:"link"`
	WalletLink nineora.Link           `json:"wallet_link"`
	KeyType    ninekey.Type           `json:"key_type"`
	NetworkID  nineora.NetworkID      `json:"network_id"`
	Password   []byte                 `json:"password"`
	Meta       map[string]interface{} `json:"meta"`
}

type AccQueryReq struct {
	Page *pagination.Page `json:"page"`
}

type BalQueryByAccReq struct {
	AccountLink nineora.Link     `json:"account_link"`
	Page        *pagination.Page `json:"page"`
}

type BillQueryByAccReq struct {
	AccountLink nineora.Link      `json:"account_link"`
	TokenIn     []nineora.TokenID `json:"token_in"`
	Page        *pagination.Page  `json:"page"`
}

type BillQueryByBalReq struct {
	BalanceID nineora.BalanceID `json:"balance_id"`
	Page      *pagination.Page  `json:"page"`
}

const (
	WalletPath         = "wallet"
	WalletBindPath     = WalletPath + "/bind"
	WalletGetPath      = WalletPath + "/g"
	AccCreatePath      = WalletPath + "/acc/create"
	AccQueryPath       = WalletPath + "/acc/q"
	AccGetPath         = WalletPath + "/acc/g"
	BalQueryByAccPath  = WalletPath + "/bal/q/by_acc"
	BillQueryByBalPath = WalletPath + "/bill/q/by_bal"
)

type WalletBind func(req *WalletBindReq) (*nineora.Wallet, *errors.Error)
type AccCreate func(req *AccCreateReq) (*nineora.Account, *errors.Error)
type AccQuery func(req *AccQueryReq) (*pagination.Pagination[nineora.Account], *errors.Error)
type BalQueryByAcc func(req *BalQueryByAccReq) (*pagination.Pagination[nineora.Balance], *errors.Error)
type BillQueryByBal func(req *BillQueryByBalReq) (*pagination.Pagination[nineora.Bill], *errors.Error)
type WalletGet func(link nineora.Link) (*nineora.Wallet, *errors.Error)
type AccGet func(link nineora.Link) (*nineora.Account, *errors.Error)
