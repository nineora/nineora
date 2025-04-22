package nineora

import (
	"github.com/nineora/nineora/nine/chain"
	"math"
)

const (
	TokenSupplyInfinite = math.MaxInt64
)

type Token struct {
	NID       TokenID       `json:"nid"`
	NetworkID NetworkID     `json:"network_id"`
	Chain     chain.Chain   `json:"chain"`
	Network   chain.Address `json:"network"`
	Address   chain.Address `json:"address"`
	Symbol    string        `json:"symbol"`
	Name      string        `json:"name"`
	Icon      string        `json:"icon"`
	Ctrl      Ctrl          `json:"ctrl"`
	Decimals  uint8         `json:"decimals"`
	Supply    uint64        `json:"supply"`
	Limit     uint64        `json:"limit"`
	UpdateAt  int64         `json:"update_at"`
	ChainAt   int64         `json:"chain_at"`
}

func (t *Token) IsInfinite() bool {
	return t.Limit == TokenSupplyInfinite
}

func (t *Token) GetSupply() Amount {
	return Amount{
		Decimals: t.Decimals,
		Amount:   t.Supply,
	}
}

func (t *Token) GetLimit() Amount {
	return Amount{
		Decimals: t.Decimals,
		Amount:   t.Limit,
	}
}
