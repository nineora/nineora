package nineora

import "github.com/nineora/nineora/nine/chain"

type Network struct {
	NID      NetworkID     `json:"nid"`
	Chain    chain.Chain   `json:"chain"`
	Address  chain.Address `json:"address"`
	Symbol   string        `json:"symbol"`
	Ctrl     Ctrl          `json:"ctrl"`
	Fee      uint64        `json:"fee"`
	UpdateAt int64         `json:"update_at"`
	ChainAt  int64         `json:"chain_at"`
}
