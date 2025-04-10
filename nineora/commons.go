package nineora

import "github.com/hootuu/gelato/strs"

type Address = string

type Ctrl uint64

func CtrlOfStr(str string) Ctrl {
	return Ctrl(strs.ToUint64(str))
}

type Chain = string

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
