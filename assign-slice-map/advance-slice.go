package main

import "fmt"

func badAdvanceSlice() {
	as := NewAdviceSlice()
	newList := as.GetBad()
	as.Update()

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}

func goodAdvanceSlice() {
	as := NewAdviceSlice()
	newList := as.GetGood()
	as.Update()

	fmt.Println("result: ")
	for _, v := range newList {
		fmt.Println(v)
	}
}

type adviceSlice struct {
	list []string
}

func NewAdviceSlice() *adviceSlice {
	return &adviceSlice{
		list: []string{
			"apple",
			"banana",
			"cherry",
		},
	}
}

func (b adviceSlice) GetBad() []string {
	return b.list
}

func (b adviceSlice) GetGood() []string {
	newList := make([]string, len(b.list))
	copy(newList, b.list)
	return newList
}

func (b adviceSlice) Update() {
	b.list[0] = "data is updated"
}
