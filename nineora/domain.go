package nineora

type DomainID = NID

type Domain struct {
	NID       DomainID `json:"nid"`
	NineoraID ID       `json:"nineora_id"`

	Chain     Chain     `json:"chain"`
	Nineora   Address   `json:"nineora"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`
}
