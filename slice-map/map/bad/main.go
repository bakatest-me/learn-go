package main

import (
	"fmt"
)

func main() {
	before := map[string]bool{
		"apple":  true,
		"banana": true,
		"cherry": true,
	}
	after := before
	before["apple"] = false

	for k, v := range after {
		fmt.Printf("%v:\t %v\n", k, v)
	}
}
