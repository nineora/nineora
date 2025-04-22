package nineora

import "github.com/nineora/nineora/nine/chain"

type Event struct {
	NID       EventID                `json:"nid"`
	Chain     chain.Chain            `json:"chain"`
	Tx        string                 `json:"tx"`
	Seq       string                 `json:"seq"`
	EventType string                 `json:"event_type"`
	CtxType   string                 `json:"ctx_type"`
	Biz       uint16                 `json:"biz"`
	Memo      string                 `json:"memo"`
	Data      map[string]interface{} `json:"data"`
	CreateAt  int64                  `json:"create_at"`
	ChainAt   int64                  `json:"chain_at"`
}
