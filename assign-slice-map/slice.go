package main

import "fmt"

func badSlice() {
	var list = []string{
		"apple",
		"banana",
		"cherry",
	}
	newList := list
	list[0] = "data is updated"

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}

func goodSlice() {
	var list = []string{
		"apple",
		"banana",
		"cherry",
	}
	newList := make([]string, len(list))
	copy(newList, list)
	list[0] = "data is updated"

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}
