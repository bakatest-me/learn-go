package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	b := true
	d := "string"

	fmt.Println("size of: struct", unsafe.Sizeof(a))
	fmt.Println("size of: bool", unsafe.Sizeof(b))
	fmt.Println("size of: string", unsafe.Sizeof(d))
}
