package daka

import "github.com/nineora/nineora/nine/nineora"

type AreaAccount struct {
	Digi nineora.Lockable `json:"digi"`
	Kaka nineora.Lockable `json:"kaka"`
	Dada nineora.Lockable `json:"dada"`
	Daka nineora.Lockable `json:"daka"`
}

type CommunityAccount struct {
	Digi nineora.Lockable `json:"digi"`
	Kaka nineora.Lockable `json:"kaka"`
	Dada nineora.Lockable `json:"dada"`
	Daka nineora.Lockable `json:"daka"`
}

type MerchantAccount struct {
	Kaka nineora.Lockable `json:"kaka"`
	Dada nineora.Lockable `json:"dada"`
	Daka nineora.Lockable `json:"daka"`
}

type MemberAccount struct {
	Kaka nineora.Lockable `json:"kaka"`
	Daka nineora.Lockable `json:"daka"`
}
