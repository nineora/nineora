package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
	"github.com/nineora/nineora/nine/chain"
)

type NID = string

func NewNID(chain chain.Chain, addr chain.Address) NID {
	idStr := fmt.Sprintf("%s@%d", addr, chain)
	return md5x.MD5(idStr)
}

type NetworkID = NID

func NewNetworkID(chain chain.Chain, addr chain.Address) NetworkID {
	return NewNID(chain, addr)
}

type NodeID = NID

func NewNodeID(chain chain.Chain, addr chain.Address) NodeID {
	return NewNID(chain, addr)
}

type TokenID = NID

func NewTokenID(chain chain.Chain, addr chain.Address) TokenID {
	return NewNID(chain, addr)
}

type BusinessID = NID

func NewBusinessID(chain chain.Chain, addr chain.Address) BusinessID {
	return NewNID(chain, addr)
}

type EventID = NID

func NewEventID(chain chain.Chain, tx string, seq string) EventID {
	return md5x.MD5(fmt.Sprintf("%s:%s@%d", tx, seq, chain))
}

type WalletID = NID

type AccountID = NID

func NewAccountID(chain chain.Chain, account chain.Address) AccountID {
	return NewNID(chain, account)
}

type BalanceID = NID

func NewBalanceID(chain chain.Chain, account chain.Address, token chain.Address) BalanceID {
	return md5x.MD5(fmt.Sprintf("%s:%s@%d", account, token, chain))
}

type BillID = NID
