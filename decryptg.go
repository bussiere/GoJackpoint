package main

import (
	"fmt"
	"os"

	"crypto/aes"

	"crypto/cipher"

	"encoding/base64"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var ciphertext []byte

func main() {

	// Load the base64 ciphertext from the arguement.

	if len(os.Args) > 1 {

		// Decode the ciphertext to put it in a usable binary format.

		dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(os.Args[1])))

		base64.StdEncoding.Decode(dbuf, []byte(os.Args[1]))

		ciphertext = []byte(dbuf)

	} else {

		fmt.Printf("Error: At least one argument required!")

		os.Exit(-1)

	}

	// Load the key from the second argument.

	key_text := "32o4908go293hohg98fh40gh"

	if len(os.Args) > 2 {

		key_text = os.Args[2]

	}

	// We chose our cipher type here in this case

	// we are using AES.

	c, err := aes.NewCipher([]byte(key_text))

	if err != nil {

		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)

		os.Exit(-1)

	}

	// We use the CFBDecrypter in order to decrypt

	// the whole stream of ciphertext using the

	// cipher setup with c and a iv.

	cfb := cipher.NewCFBDecrypter(c, commonIV)

	plaintext := make([]byte, len(ciphertext))

	cfb.XORKeyStream(plaintext, ciphertext)

	// We then print out the resulting text.

	fmt.Printf("%x=>%s\n", ciphertext, plaintext)

}
