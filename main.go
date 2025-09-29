package main

import (
	"cmp"
	"fmt"
)

func main() {
	code := 200
	val := cmp.Or(0, map[bool]int{true: 50}[code == 200])
	fmt.Println("[x] print", val)
}
