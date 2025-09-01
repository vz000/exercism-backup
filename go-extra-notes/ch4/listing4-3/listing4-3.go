package main

// This program creates a deadlock
// A dead lock is any situtation in which no member of some group of entities
// can proceed because each watis for another member, including itself.

import (
	"fmt"
)

func main() {
	// Declare an unbuffered channel
	counter := make(chan int)
	// PPerform send operation by launching new goroutine
	go func() {
		counter <- 10
	}()

	fmt.Println(<-counter) //Receive operation from channel
}
