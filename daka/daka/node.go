package daka

type Unimkt struct {
	Lotto Lotto `json:"lotto"`
}

type Zone struct {
	Account Account `json:"account"`
}

type Branch struct {
	Account Account `json:"account"`
}

type Division struct {
	Account Account `json:"account"`
}

type Associate struct {
	Account Account `json:"account"`
}

type Merchant struct {
	Account Account `json:"account"`
}

type Member struct {
	Account Account `json:"account"`
}
