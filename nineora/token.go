package nineora

type TokenID = NID
type Token struct {
	NID       TokenID  `json:"nid"`
	NineoraID ID       `json:"nineora_id"`
	DomainID  DomainID `json:"domain_id"`

	Nineora Address `json:"nineora"`
	Domain  Address `json:"domain"`

	Chain     Chain     `json:"chain"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`
}

type TokenCore[T any] struct {
	Token T `json:"token"`
}
