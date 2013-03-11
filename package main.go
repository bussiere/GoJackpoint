package main

import (
	"code.google.com/p/go.crypto/openpgp"
	"code.google.com/p/go.crypto/openpgp/armor"
	"fmt"
	"os"
)

func main() {
	w, _ := armor.Encode(os.Stdout, "PGP MESSAGE", nil)
	plaintext, _ := openpgp.SymmetricallyEncrypt(w, []byte("golang"), nil)
	fmt.Fprintf(plaintext, "Hello from golang.\n")
	plaintext.Close()
	w.Close()
	fmt.Print("\n")
}
