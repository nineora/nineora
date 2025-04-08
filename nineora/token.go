package nineora

type TokenID = NID
type Token struct {
	NID       TokenID   `json:"nid"`
	NineoraID ID        `json:"nineora_id"`
	NetworkID NetworkID `json:"network_id"`

	Nineora Address `json:"nineora"`
	Network Address `json:"network"`

	Chain     Chain     `json:"chain"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	Supply uint64 `json:"supply"`
	Limit  uint64 `json:"limit"`

	TokenType string                 `json:"token_type"`
	Token     map[string]interface{} `json:"token"`
}

type TokenCore[T any] struct {
	Token T `json:"token"`
}
