package chain

type Address = string

type Pair struct {
	Chain   Chain   `json:"chain"`
	Address Address `json:"address"`
}
