package nineora

type AccountID = NID
type Account struct {
	Address Address
	Owner   Address
	Token   Address
	Balance uint64
}
