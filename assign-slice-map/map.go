package main

import (
	"fmt"
	"maps"
)

func badMap() {
	var oldMap = map[string]string{
		"apple":  "apple",
		"banana": "banana",
		"cherry": "cherry",
	}
	newMap := oldMap
	oldMap["apple"] = "data is updated"

	fmt.Println("result: ")
	for _, v := range newMap {
		fmt.Println(v)
	}
}

func goodMap() {
	var oldMap = map[string]string{
		"apple":  "apple",
		"banana": "banana",
		"cherry": "cherry",
	}
	newMap := maps.Clone(oldMap)
	oldMap["apple"] = "data is updated"

	fmt.Println("result: ")
	for _, v := range newMap {
		fmt.Println(v)
	}
}
