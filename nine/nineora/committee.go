package nineora

import "github.com/nineora/nineora/nine/chain"

type Member struct {
	Weight     uint64 `json:"weight"`
	Role       uint64 `json:"role"`
	JoinedTime uint64 `json:"joined_time"`
	Nine       bool   `json:"nine"`
}

type Committee struct {
	Members map[chain.Address]Member `json:"members,omitempty"`
}
