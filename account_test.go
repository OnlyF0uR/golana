package golana_test

import (
	"testing"

	"github.com/OnlyF0uR/golana"
)

func TestAccount(t *testing.T) {
	// Create a new account
	acc, err := golana.CreateAccount()
	if err != nil {
		t.Fatal(err)
	}

	// Import the new account
	acc2, err := golana.AccountFromPrivateKeyBase58(acc.PrivateKey())
	if err != nil {
		t.Fatal(err)
	}

	// Check if the public keys match
	if acc.PublicKey() != acc2.PublicKey() {
		t.Fatal("Public keys do not match")
	}

	// Log the keys
	t.Logf("Public key: %s", acc2.PublicKey())
	t.Logf("Private key: %s", acc2.PrivateKey())
}

func TestSignMessage(t *testing.T) {
	// Create account
	acc, err := golana.CreateAccount()
	if err != nil {
		t.Fatal(err)
	}

	// Sign a message
	msg := []byte("Hello, world!")
	sig, err := acc.SignMessage(msg)
	if err != nil {
		t.Fatal(err)
	}

	// Log the signature
	t.Logf("Signature: %s", sig)

	// Verify the signature
	ok, err := acc.VerifySignature(msg, sig)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("Signature was not a match")
	}
}
