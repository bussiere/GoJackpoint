package main

import (
	"bytes"
	"code.google.com/p/go.crypto/openpgp"
	"encoding/base64"
	"fmt"
)

func main() {

	//openpgp.NewEntity("bussiere", "test", "bussiere@gmail.com", nil)

	var entity *openpgp.Entity
	entity, err := openpgp.NewEntity("bussiere", "test", "bussiere@gmail.com", nil)
	if err != nil {

	}

	var (
		buffer bytes.Buffer
	)

	entity.SerializePrivate(&buffer, nil)
	data := base64.StdEncoding.EncodeToString([]byte(buffer.String()))

	fmt.Printf("%q\n", data)

	entity.Serialize(&buffer)
	data2 := base64.StdEncoding.EncodeToString([]byte(buffer.String()))

	fmt.Printf("%q\n", data2)

	//fmt.Printf(buffer.String())
}
