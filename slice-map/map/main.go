package main

import "fmt"

func main() {
	m := map[string]bool{
		"apple":  true,
		"banana": true,
		"cherry": true,
	}

	for k, v := range m {
		fmt.Printf("%v:\t %v\n", k, v)
	}
}
