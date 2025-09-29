package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------- bad slice ----------")
	badSlice()
	fmt.Println("---------- good slice ----------")
	goodSlice()
	fmt.Println()

	fmt.Println("---------- bad map ----------")
	badMap()
	fmt.Println("---------- good map ----------")
	goodMap()
	fmt.Println()

	fmt.Println("-------------- bad advance slice --------------")
	badAdvanceSlice()
	fmt.Println("-------------- good advance slice --------------")
	goodAdvanceSlice()

	fmt.Println("-------------- bad advance map --------------")
	badAdvanceMap()
	fmt.Println("-------------- good advance map --------------")
	goodAdvanceMap()
}
