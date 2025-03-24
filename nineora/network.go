package nineora

type Network[T any] struct {
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Domain    Address   `json:"domain"`
	Owner     Address   `json:"owner"`
	Committee Committee `json:"committee"`
	Network   T         `json:"network"`
	Collar    Address   `json:"collar"`
}

type Node[T any] struct {
	Address  Address `json:"address"`
	Network  Address `json:"network"`
	Superior Address `json:"superior"`
	Owner    Address `json:"owner"`
	Benefit  Address `json:"benefit"`
	Ctrl     Ctrl    `json:"ctrl"`
	Node     T       `json:"node"`
}
