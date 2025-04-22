package ninekey

import (
	"github.com/hootuu/gelato/errors"
)

type Type int

const (
	Ed25519   Type = 1
	Secp256k1 Type = 2
)

func (t Type) String() string {
	switch t {
	case Ed25519:
		return "Ed25519"
	case Secp256k1:
		return "Secp256k1"
	}
	return ""
}

func (t Type) Verify() *errors.Error {
	switch t {
	case Ed25519, Secp256k1:
		return nil
	}
	return errors.Verify("invalid key type")
}
