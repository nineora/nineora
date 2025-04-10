package nineora

type BusinessID = NID
type Business struct {
	NID       BusinessID `json:"nid"`
	NineoraID ID         `json:"nineora_id"`
	NetworkID NetworkID  `json:"network_id"`

	Chain     Chain     `json:"chain"`
	Nineora   Address   `json:"nineora"`
	Network   Address   `json:"domain"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	BizType string                 `json:"biz_type"`
	Biz     map[string]interface{} `json:"biz"`

	TimestampMsUTC int64 `json:"timestamp_ms_utc"`
}

type Biz[T any] struct {
	Biz T `json:"biz"`
}
