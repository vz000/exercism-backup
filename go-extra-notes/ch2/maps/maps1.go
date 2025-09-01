package main

import (
	"fmt"
	"sort"
)

// This is listing 2-15
// Iterates over a map with an order
// To specify the order, this example maintain a slice to store keys of a map
func listing2_15() {
	// Initialize map with make function
	chapts := make(map[int]string)

	// Add data as key/value pairs
	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"

	// Slice for specifying the order or the map
	var keys []int
	// Appending keys of the map
	for k := range chapts {
		keys = append(keys, k)
	}

	// Ints sorts a slice of ints in increasing order.
	sort.Ints(keys)
	// Iterate over the map with an order
	for _, k := range keys {
		fmt.Println("Key: ", k, "Value: ", chapts[k])
	}
}

func main() {
	// Declares a nil map
	// map[KeyType]ValueType
	// This is similar as Sets in Redis (unique values but order may not be preserved)
	var chapts map[int]string

	// Initialize map with make function
	chapts = make(map[int]string)

	// Add data as key/value pairs
	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"

	// Iterate over the elements of map using range
	for k, v := range chapts {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}

	// Declare and initialize map using map literal
	langs := map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}

	// Delete an element
	delete(langs, "EL")

	// Lookout an element with key
	if lan, ok := langs["EL"]; ok {
		fmt.Println(lan)
	} else {
		fmt.Println("\nKey doesn't exist\n")
	}

	fmt.Println("Listing 2-15:\n")
	listing2_15()
}
