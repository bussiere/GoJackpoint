package main

import (
	"./structure"
	"fmt"
)

func main() {
	j := new(struct_jackpoint.Jack)
	j.Id = 2
	fmt.Printf("%d", j.Id)

}
