package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/collection"
	"github.com/hootuu/gelato/crtpto/md5x"
)

type AccountID = NID

func NewAccountID(chain Chain, account Address) AccountID {
	return md5x.MD5(fmt.Sprintf("[%s]%s", chain, account))
}

type Account struct {
	NID            AccountID `json:"nid"`
	Chain          Chain     `json:"chain"`
	Address        Address   `json:"address"`
	TimestampMsUTC int64     `json:"timestamp_ms_utc"`

	WithBalance map[string]*BalanceItem `json:"with_balance"`
}

type BalanceID = NID

func NewBalanceID(chain Chain, account Address, token Address) BalanceID {
	return md5x.MD5(fmt.Sprintf("[%s]%s:%s", chain, account, token))
}

type Balance struct {
	NID            BalanceID `json:"nid"`
	AccountID      AccountID `json:"account_id"`
	TokenID        TokenID   `json:"token_id"`
	Chain          Chain     `json:"chain"`
	Token          Address   `json:"token"`
	Account        Address   `json:"account"`
	Balance        uint64    `json:"balance"`
	TimestampMsUTC int64     `json:"timestamp_ms_utc"`
}

type BalanceItem struct {
	TokenID TokenID `json:"token_id"`
	Symbol  string  `json:"symbol"`
	Amount  Amount  `json:"amount"`
}

type Direction int8

const (
	IN  Direction = 1
	OUT Direction = -1
)

type BillID = NID

func NewBillID(chain Chain, address Address, direction Direction, event NID) BillID {
	return md5x.MD5(fmt.Sprintf("[%s]%s[%d]%s", chain, address, direction, event))
}

type Bill struct {
	NID       BillID          `json:"nid"`
	AccountID AccountID       `json:"account_id"`
	BalanceID BalanceID       `json:"balance_id"`
	Direction Direction       `json:"direction"`
	Sender    NID             `json:"sender"`
	Recipient NID             `json:"recipient"`
	Amount    Amount          `json:"amount"`
	Biz       string          `json:"biz"`
	Ctx       collection.Dict `json:"ctx"`
}
