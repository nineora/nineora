package nineora

import (
	"github.com/nineora/nineora/nine/chain"
	"github.com/stretchr/objx"
)

type Node struct {
	NID      NodeID                 `json:"nid"`
	Chain    chain.Chain            `json:"chain"`
	Network  chain.Address          `json:"network"`
	Deep     uint32                 `json:"deep"`
	Address  chain.Address          `json:"address"`
	Superior chain.Address          `json:"superior"`
	Benefit  chain.Address          `json:"benefit"`
	Ctrl     Ctrl                   `json:"ctrl"`
	Node     map[string]interface{} `json:"node"`
	UpdateAt int64                  `json:"update_at"`
	ChainAt  int64                  `json:"chain_at"`
}

func (n *Node) NodeX() objx.Map {
	return objx.New(n.Node)
}
