package nineora

type Member struct {
	Weight     uint64 `json:"weight"`
	Role       uint64 `json:"role"`
	JoinedTime uint64 `json:"joined_time"`
	Nine       bool   `json:"nine"`
}

type Committee struct {
	Members map[Address]Member `json:"members,omitempty"`
}
