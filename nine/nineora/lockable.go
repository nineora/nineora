package nineora

type Lockable struct {
	Balance uint64 `json:"balance"`
	Locked  uint64 `json:"locked"`
}

func (l Lockable) TotalBalance() uint64 {
	return l.Balance + l.Locked
}

func (l Lockable) AvailableBalance() uint64 {
	return l.Balance
}
