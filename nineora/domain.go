package nineora

type Domain struct {
	Address Address `json:"address"`
	Symbol  string  `json:"symbol"`
	Owner   Address `json:"owner"`
	Collar  Address `json:"collar"`
}
