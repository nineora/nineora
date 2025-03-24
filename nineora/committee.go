package nineora

type Member struct {
	Weight     uint64
	Role       uint64
	JoinedTime uint64
	Nine       bool
}

type Committee struct {
	Members map[Address]Member
}
