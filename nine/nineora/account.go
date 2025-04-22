package nineora

//
//import (
//	"fmt"
//	"github.com/hootuu/gelato/crtpto/md5x"
//	"github.com/nineora/nineora/nine/chain"
//	"github.com/stretchr/objx"
//)
//
//type Account struct {
//	NID      AccountID     `json:"nid"`
//	Chain    chain.Chain   `json:"chain"`
//	Address  chain.Address `json:"address"`
//	UpdateAt int64         `json:"update_at"`
//	ChainAt  int64         `json:"chain_at"`
//	Balance  []*Balance    `json:"balance"`
//}
//
//type Balance struct {
//	NID       BalanceID `json:"nid"`
//	AccountID AccountID `json:"account_id"`
//	TokenID   TokenID   `json:"token_id"`
//	Symbol    string    `json:"symbol"`
//	Name      string    `json:"name"`
//	Icon      string    `json:"icon"`
//	Decimals  uint8     `json:"decimals"`
//	Balance   uint64    `json:"balance"`
//	UpdateAt  int64     `json:"update_at"`
//	ChainAt   int64     `json:"chain_at"`
//}
//
//func (bal *Balance) GetBalance() Amount {
//	return Amount{
//		Decimals: bal.Decimals,
//		Amount:   bal.Balance,
//	}
//}
//
//type Direction int8
//
//const (
//	IN  Direction = 1
//	OUT Direction = -1
//)
//
//type BillID = NID
//
//func NewBillID(chain chain.Chain, address chain.Address, direction Direction, event NID) BillID {
//	return md5x.MD5(fmt.Sprintf("%s:%d:%s@%s", address, direction, event, chain))
//}
//
//type Bill struct {
//	NID       BillID                 `json:"nid"`
//	AccountID AccountID              `json:"account_id"`
//	BalanceID BalanceID              `json:"balance_id"`
//	TokenID   TokenID                `json:"token_id"`
//	Direction Direction              `json:"direction"`
//	Sender    chain.Address          `json:"sender"`
//	Recipient chain.Address          `json:"recipient"`
//	Symbol    string                 `json:"symbol"`
//	Decimals  uint8                  `json:"decimals"`
//	Amount    uint64                 `json:"amount"`
//	Balance   uint64                 `json:"balance"`
//	Biz       string                 `json:"biz"`
//	Ctx       map[string]interface{} `json:"ctx"`
//}
//
//func (b *Bill) CtxX() objx.Map {
//	return objx.New(b.Ctx)
//}
//todo del
