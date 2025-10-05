package main

import (
	"fmt"
)

func main() {
	before := []string{
		"apple",
		"banana",
		"cherry",
	}
	fmt.Println("Before: ", before)
	/* Before:  [apple banana cherry] */
	after := before
	before[0] = "__imposter__"

	fmt.Println("After: ", after)
	/* After:  [__imposter__ banana cherry] */
}
