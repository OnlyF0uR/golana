package solcrypt

import (
	"crypto"
	"crypto/ed25519"
	crypto_rand "crypto/rand"
	"errors"
)

func GeneratePair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(crypto_rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return pubKey, privKey, nil
}

func FromPrivateKey(privateKey []byte) (ed25519.PublicKey, ed25519.PrivateKey) {
	privKey := ed25519.PrivateKey(privateKey)
	return privKey.Public().(ed25519.PublicKey), privKey
}

func Sign(privKey ed25519.PrivateKey, payload []byte) ([]byte, error) {
	digest, err := privKey.Sign(crypto_rand.Reader, payload, crypto.Hash(0))
	if err != nil {
		return nil, err
	}

	return digest, nil
}

func VerifySig(publicKey ed25519.PublicKey, payload, signature []byte) (bool, error) {
	if len(signature) != ed25519.SignatureSize {
		return false, errors.New("invalid signature length")
	}

	return ed25519.Verify(publicKey, payload, signature), nil
}
