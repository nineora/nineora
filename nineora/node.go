package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
)

type NodeID = NID

func NewNodeID(chain Chain, addr Address) NodeID {
	return md5x.MD5(fmt.Sprintf("[%s]%s", chain, addr))
}

type Node struct {
	NID     NodeID  `json:"nid"`
	Chain   Chain   `json:"chain"`
	Nineora Address `json:"nineora"`
	Network Address `json:"network"`

	Address  Address `json:"address"`
	Superior Address `json:"superior"`
	Owner    Address `json:"owner"`
	Benefit  Address `json:"benefit"`
	Ctrl     Ctrl    `json:"ctrl"`

	NodeType string                 `json:"node_type"`
	Node     map[string]interface{} `json:"node"`

	TimestampMsUTC int64 `json:"timestamp_ms_utc"`
}

type NodeCore[T any] struct {
	Node T `json:"node"`
}
