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
	Collar    Address   `json:"collar"`
	Committee Committee `json:"committee"`

	Treasury Address `json:"treasury"`
	Meta     Address `json:"meta"`
	Supply   uint64  `json:"supply"`
	Limit    uint64  `json:"limit"`

	TokenType string `json:"token_type"`

	TimestampMsUTC int64 `json:"timestamp_ms_utc"`
}

type TokenCore[T any] struct {
	Token T `json:"token"`
}
