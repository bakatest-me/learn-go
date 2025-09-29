package main

import (
	"errors"
	"fmt"
)

var ErrSaveOrder = errors.New("error save order")
var ErrOrderNotFound = errors.New("order not found")
var ErrProductNotFound = errors.New("product not found")

func main() {
	e := fmt.Errorf("%w, %w", ErrOrderNotFound, ErrSaveOrder)
	fmt.Println(e)
	fmt.Println("errors.Is(e, ErrOrderNotFound): ", errors.Is(e, ErrOrderNotFound))
	fmt.Println("errors.Is(e, ErrSaveOrder): ", errors.Is(e, ErrSaveOrder))
	fmt.Println()

	we := WrapErr(ErrOrderNotFound, ErrSaveOrder)
	fmt.Println(we)
	fmt.Println("errors.Is(e, ErrOrderNotFound): ", errors.Is(we, ErrOrderNotFound))
	fmt.Println("errors.Is(e, ErrSaveOrder): ", errors.Is(we, ErrSaveOrder))
	fmt.Println()
	// wrap in wrap
	we2 := WrapErr(we, ErrProductNotFound)
	fmt.Println(we2)
	fmt.Println("errors.Is(e, ErrOrderNotFound): ", errors.Is(we2, ErrOrderNotFound))
	fmt.Println("errors.Is(e, ErrSaveOrder): ", errors.Is(we2, ErrSaveOrder))
	fmt.Println("errors.Is(e, ErrProductNotFound): ", errors.Is(we2, ErrProductNotFound))
}

func WrapErr(err ...error) error {
	if len(err) == 0 {
		return nil
	}
	if len(err) == 1 {
		return err[0]
	}

	var (
		format    = ""
		lastIndex = len(err) - 1
		args      = make([]any, len(err))
	)

	for i, e := range err {
		args[i] = e
		if i == lastIndex {
			format += "%w"
			continue
		}
		format += "%w, "
	}
	return fmt.Errorf(format, args...)
}
