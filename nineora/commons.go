package nineora

type Address string

type Ctrl uint64

type Chain string

const (
	SOLANA Chain = "SOLANA"
	SUI    Chain = "SUI"
)

func SupportedChains() []Chain {
	return []Chain{
		SOLANA,
		SUI,
	}
}
