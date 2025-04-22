package chain

type Chain int

const (
	Nineora Chain = 0
	SUI     Chain = 1
	SOLANA  Chain = 2
)

func (chain Chain) Name() string {
	switch chain {
	case Nineora:
		return "nineora"
	case SUI:
		return "sui"
	case SOLANA:
		return "solana"
	}
	return ""
}

func Support() []Chain {
	return []Chain{Nineora, SUI, SOLANA}
}
