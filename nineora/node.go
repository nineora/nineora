package nineora

type NodeID = NID
type Node struct {
	NID       NodeID    `json:"nid"`
	NineoraID ID        `json:"nineora_id"`
	DomainID  DomainID  `json:"domain_id"`
	NetworkID NetworkID `json:"network_id"`

	Nineora Address `json:"nineora"`
	Domain  Address `json:"domain"`

	Chain    Chain   `json:"chain"`
	Address  Address `json:"address"`
	Network  Address `json:"network"`
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
