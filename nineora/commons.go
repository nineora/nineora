package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/strs"
	"strconv"
	"strings"
)

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

type Amount struct {
	Decimals uint8  `json:"decimals"`
	Amount   uint64 `json:"amount"`
}

func (a Amount) Display() string {
	balanceStr := strconv.FormatUint(a.Amount, 10)
	d := int(a.Decimals)
	if d == 0 {
		return balanceStr
	}

	var integerPart, fractionalPart string

	if len(balanceStr) > d {
		integerPart = balanceStr[:len(balanceStr)-d]
		fractionalPart = balanceStr[len(balanceStr)-d:]
	} else {
		padded := strings.Repeat("0", d-len(balanceStr)) + balanceStr
		integerPart = "0"
		fractionalPart = padded
	}

	integerPart = trimLeadingZeros(integerPart)

	return fmt.Sprintf("%s.%s", integerPart, fractionalPart)
}

func trimLeadingZeros(s string) string {
	trimmed := strings.TrimLeft(s, "0")
	if trimmed == "" {
		return "0"
	}
	return trimmed
}
