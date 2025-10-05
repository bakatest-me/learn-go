package main

import "fmt"

func main() {
	before := []string{
		"apple",
		"banana",
		"cherry",
	}
	fmt.Println("before: ", before)

	after := before
	UpdateSlice(before)

	fmt.Println("after: ", after)
}

func UpdateSlice(v []string) {
	if len(v) <= 0 {
		return
	}
	v[0] = "__imposter__"
}
