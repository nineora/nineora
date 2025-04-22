package ninerpc

// todo
//import (
//	"github.com/hootuu/gelato/errors"
//	"github.com/hootuu/gelato/io/pagination"
//	"github.com/nineora/nineora/nine/nineora"
//)
//
//type AccountQueryReq struct {
//	WithBalance bool             `json:"with_balance"`
//	Page        *pagination.Page `json:"page"`
//}
//
//type AccountGetReq struct {
//	NID         nineora.AccountID `json:"nid"`
//	WithBalance bool              `json:"with_balance"`
//}
//
//type BillQueryReq struct {
//	AccountID nineora.AccountID `json:"account_id"`
//	Token     nineora.TokenID   `json:"token"`
//	Page      *pagination.Page  `json:"page"`
//}
//
//const (
//	AccountPath      = "account"
//	AccountQueryPath = AccountPath + "/q"
//	AccountGetPath   = AccountPath + "/get"
//	BillQueryPath    = AccountPath + "/bill/q"
//)
//
//type AccountQuery func(req *AccountQueryReq) (*pagination.Pagination[nineora.Account], *errors.Error)
//type AccountGet func(req *AccountGetReq) (*nineora.Account, *errors.Error)
//type BillQuery func(req *BillQueryReq) (*pagination.Pagination[nineora.Bill], *errors.Error)
