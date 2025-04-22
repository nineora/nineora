package ninekey

import (
	"encoding/hex"
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/mr-tron/base58/base58"
	"github.com/nineora/nineora/nine/chain"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

type PrivateKey []byte
type PublicKey []byte
type Mnemonic []byte

type Key struct {
	Type       Type       `json:"type"`
	PrivateKey PrivateKey `json:"-"`
	Mnemonic   Mnemonic   `json:"-"`
	PublicKey  PublicKey  `json:"public_key"`
}

type AddressGet func(key PublicKey) (chain.Address, *errors.Error)

var gAddressFactory map[chain.Chain]AddressGet

func init() {
	gAddressFactory = make(map[chain.Chain]AddressGet)
	gAddressFactory[chain.Nineora] = func(key PublicKey) (chain.Address, *errors.Error) {
		schemedPubKey := append([]byte{0x00}, key...)
		hash := sha3.Sum256(schemedPubKey)
		addr := base58.EncodeAlphabet(hash[:], base58.BTCAlphabet)
		return chain.Address(addr), nil
	}
	gAddressFactory[chain.SUI] = func(key PublicKey) (chain.Address, *errors.Error) {
		schemedPubKey := append([]byte{0x00}, key...)
		addrBytes := blake2b.Sum256(schemedPubKey)
		return chain.Address(fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])), nil
	}
	gAddressFactory[chain.SOLANA] = func(key PublicKey) (chain.Address, *errors.Error) {
		addr := base58.EncodeAlphabet(key, base58.BTCAlphabet)
		return chain.Address(addr), nil
	}
}

func CalcAddress(chain chain.Chain, key PublicKey) (chain.Address, *errors.Error) {
	addrGet, ok := gAddressFactory[chain]
	if !ok {
		return "", errors.Verify("no such address getter for: " + chain.Name())
	}
	return addrGet(key)
}

func (key *Key) Address(chain chain.Chain) (chain.Address, *errors.Error) {
	return CalcAddress(chain, key.PublicKey)
}

func (pub PublicKey) Address(chain chain.Chain) (chain.Address, *errors.Error) {
	return CalcAddress(chain, pub)
}
