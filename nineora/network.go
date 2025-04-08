package nineora

type NetworkID = NID
type Network struct {
	NID       NetworkID `json:"nid"`
	NineoraID ID        `json:"nineora_id"`

	Chain   Chain   `json:"chain"`
	Nineora Address `json:"nineora"`

	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	Fee uint64 `json:"fee"`
}
