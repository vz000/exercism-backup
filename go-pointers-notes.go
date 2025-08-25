// using the & operator the memory address can be accessed
package main
import "fmt"

// Assigns the memory address of a variable to a pointer variable
func example1() {
	var name = "John"
	var ptr* string

	// assign the memory address of name to the pointer
	ptr = &name

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Address of the variable ", &name)
}

// This function example gets the value pointer by a pointer
func example2() {
	var name = "John"
	// Here * is used to declare a pointer variable
	var ptr *string

	ptr = &name

	// Use * to get the value pointed by ptr
	// here * is the dereference operator.
	// The dereference operator operates on a pointer and 
	// gives the value stored in that pointer
	fmt.Println(*ptr)
}

func example3() {
	var num int // Not initialized, has NO value
	var ptr *int // Pointing to nothing. Doesn't point to any address.

	num = 22 // now has value 22. variable ptr still points to nothing
	fmt.Println("Address of num: ", &num)
	fmt.Println("Value of num: ", num)

	ptr = &num // now points to the same address num has
	fmt.Println("\nAddress of pointer ptr: ", ptr)
	fmt.Println("Content of pointer ptr: ", *ptr)

	num = 11
	fmt.Println("\nAddress of pointer ptr: ", ptr)
	fmt.Println("Content of pointer ptr: ", *ptr)

	*ptr = 2
	fmt.Println("\nAddress of num: ", &num)
	fmt.Println("Value of num: ", num)

	// Note: pointers can be created without using the * operator
	var name1 = "John"
	var ptr1 = &name1

	fmt.Println("Value of ptr: ", ptr1)
	fmt.Println("Address of name: ", &name1)
}

// We can pass pointers as arguments to a function
func example4(num *int) {
	// dereference the pointer
	*num = 30
}

// Pointers can be returned from functions
func example5() *string {
	message := "Example5"

	// returns the address of message
	return &message
}

func main() {
	var num int = 5
	// print the value stored in variable
	fmt.Println("Variable value: ", num)

	// prints the address of the variable
	fmt.Println("Memory address: ", &num)

	// use the pointer variables to store the memory address
	//var num2 int = 6

	// *int represents that the pointer variable is of int type.
	//		in other words, stores the memory address of int variable
	//var ptr1 *int = &num // stores the memory address of the num variable

	// Create pointer variables of other types.
	//var ptr2 *string

	// calling example 1
	example1()

	// calling example 2
	example2()

	// Calling example 3
	fmt.Println("Output of example 3 ----")
	example3()

	// Passing pointers as arguments to functions
	fmt.Println("\nOutput produced by example 4 -----")
	var number = 55
	//function call
	example4(&number)

	fmt.Println("The number is: ", number)

	// Example 5 where a pointer is returned
	fmt.Println("\nOutput from example 5 -----")
	result := example5()
	fmt.Println("This is", *result)
}
