package nineora

type NetworkID = NID
type Network struct {
	NID       NetworkID `json:"nid"`
	NineoraID ID        `json:"nineora_id"`
	DomainID  DomainID  `json:"domain_id"`

	Chain     Chain     `json:"chain"`
	Nineora   Address   `json:"nineora"`
	Domain    Address   `json:"domain"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	NetworkType string                 `json:"network_type"`
	Network     map[string]interface{} `json:"network"`
}

type NetworkCore[T any] struct {
	Network T `json:"network"`
}
