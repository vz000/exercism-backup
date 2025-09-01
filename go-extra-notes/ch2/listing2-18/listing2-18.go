// Example that demonstrates recover
package main

import (
	"fmt"
)

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division error: %v", r)
		}
	}()

	result = a / b
	return result, nil
}

func panicRecover() {
	defer fmt.Println("Deferred call - 1")
	defer func() {
		fmt.Println("Deferred cal - 2")
		if e := recover(); e != nil {
			// e is a value passed to panic
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just panicking for the sake of example")
	fmt.Println("this will never be called")
}

func main() {
	fmt.Println("Starting to panic")
	panicRecover()
	fmt.Println("Program regains control after the panic recovery")

	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}
