package nineora

import (
	"github.com/nineora/nineora/nine/chain"
	"github.com/stretchr/objx"
)

type Business struct {
	NID       BusinessID             `json:"nid"`
	NetworkID NetworkID              `json:"network_id"`
	Chain     chain.Chain            `json:"chain"`
	Network   chain.Address          `json:"domain"`
	Address   chain.Address          `json:"address"`
	Symbol    string                 `json:"symbol"`
	Ctrl      Ctrl                   `json:"ctrl"`
	Biz       map[string]interface{} `json:"biz"`
	UpdateAt  int64                  `json:"update_at"`
	ChainAt   int64                  `json:"chain_at"`
}

func (biz *Business) BizX() objx.Map {
	return objx.New(biz.Biz)
}
