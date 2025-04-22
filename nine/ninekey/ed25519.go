package ninekey

import (
	"crypto/ed25519"
	"crypto/rand"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/logger"
	"github.com/tyler-smith/go-bip39"
	"go.uber.org/zap"
)

func Ed25519Generate() (*Key, *errors.Error) {
	entropy := make([]byte, 32)
	if _, err := rand.Read(entropy); err != nil {
		logger.Error.Error("rand.Read(entropy) err", zap.Error(err))
		return nil, errors.System("ed25519 rand read failed", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if _, err := rand.Read(entropy); err != nil {
		logger.Error.Error("bip39.NewMnemonic(entropy) err", zap.Error(err))
		return nil, errors.System("new mnemonic failed", err)
	}

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		logger.Error.Error("ed25519.GenerateKey(rand.Reader) err", zap.Error(err))
		return nil, errors.System("ed25519 generate key failed", err)
	}

	return &Key{
		Type:       Ed25519,
		PrivateKey: privateKey.Seed(),
		Mnemonic:   []byte(mnemonic),
		PublicKey:  PublicKey(publicKey),
	}, nil
}
