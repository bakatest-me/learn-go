package main

import (
	"fmt"
	"maps"
)

func badAdvanceMap() {
	as := NewAdviceMap()
	newList := as.GetBad()
	as.Update()

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}

func goodAdvanceMap() {
	as := NewAdviceMap()
	newList := as.GetGood()
	as.Update()

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}

type adviceMap struct {
	mapVal map[string]string
}

func NewAdviceMap() *adviceMap {
	return &adviceMap{
		mapVal: map[string]string{
			"apple":  "apple",
			"banana": "banana",
			"cherry": "cherry",
		},
	}
}

func (b adviceMap) GetBad() map[string]string {
	return b.mapVal
}

func (b adviceMap) GetGood() map[string]string {
	return maps.Clone(b.mapVal)
}

func (b adviceMap) Update() {
	b.mapVal["apple"] = "data is updated"
}
