package nineora

type Token[T any] struct {
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Domain    Address   `json:"domain"`
	Committee Committee `json:"committee"`
	Token     T         `json:"token"`
	Ctrl      Ctrl      `json:"ctrl"`
}
