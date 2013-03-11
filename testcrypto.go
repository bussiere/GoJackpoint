package main

import (
	"code.google.com/p/go.crypto/openpgp"
	"fmt"
)

func main() {

	//openpgp.NewEntity("bussiere", "test", "bussiere@gmail.com", nil)

	var test = openpgp.NewEntity("bussiere", "test", "bussiere@gmail.com", nil)
	test.SerializePrivate(w, config)

	fmt.Printf()
}
