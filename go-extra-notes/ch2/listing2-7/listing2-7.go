// This program demonstrates passing an anonymous function as an argument
// to another (non anonymous) function. The anonymous function closes over the variables
// to form closure.

package main

import (
	"fmt"
)

// "Accept as an argument a function with this signature"
// Signature of expected function as an arg,,,
func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y)

	x, y = f(50)
	fmt.Println(x, y)
}

func main() {
	a, b := 5, 8
	// Anonymous function
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	// Passing function value as an argument
	SplitValues(fn)

	// Calling the function value by providing an argument
	x, y := fn(20)
	fmt.Println(x,y)
}
