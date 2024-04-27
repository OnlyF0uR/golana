package golana

import (
	"crypto/ed25519"

	"github.com/OnlyF0uR/golana/solcrypt"
	"github.com/mr-tron/base58"
)

type Account struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func (a *Account) PublicKey() string {
	return base58.Encode(a.publicKey)
}

func (a *Account) PrivateKey() string {
	return base58.Encode(a.privateKey)
}

func Create() (*Account, error) {
	pubKey, privKey, err := solcrypt.GeneratePair()
	if err != nil {
		return nil, err
	}

	return &Account{privateKey: privKey, publicKey: pubKey}, nil
}

func AccountFromPrivateKey(privateKey []byte) *Account {
	pubKey, privKey := solcrypt.FromPrivateKey(privateKey)
	return &Account{privateKey: privKey, publicKey: pubKey}
}

func AccountFromPrivateKeyBase58(privateKey string) (*Account, error) {
	key, err := base58.Decode(privateKey)
	if err != nil {
		return nil, err
	}

	return AccountFromPrivateKey(key), nil
}
