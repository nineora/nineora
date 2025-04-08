package nineora

type NodeID = NID
type Node struct {
	NID        NodeID    `json:"nid"`
	NineoraID  ID        `json:"nineora_id"`
	NetworkID  NetworkID `json:"network_id"`
	SuperiorID NodeID    `json:"superior_id"`

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
}

type NodeCore[T any] struct {
	Node T `json:"node"`
}
