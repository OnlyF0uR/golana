# golana

Solana SDK library written in Go

```go
package main

import (
  "fmt"

  "github.com/OnlyF0uR/golana"
)

func main() {
  // Create account
	acc, err := golana.CreateAccount()
	if err != nil {
		panic(err)
	}

	// Sign a message
	msg := []byte("Hello, world!")
	sig, err := acc.SignMessage(msg)
	if err != nil {
		panic(err)
	}

	// Log the signature
	fmt.Printf("Signature: %s\n", sig)
}
```
