package nineora

type ID = NID
type Nineora struct {
	NID       ID        `json:"nid"`
	Chain     Chain     `json:"chain"`
	Address   Address   `json:"address"`
	Committee Committee `json:"committee"`
}
