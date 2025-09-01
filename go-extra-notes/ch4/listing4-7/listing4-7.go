package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type fibvalue struct {
	input, value int
}

var wg sync.WaitGroup

// Generated random values
// this function if the first stago of the pipeline
// First stage randomly generates values with an upprer limit of 50
// It has an outbound channel to give inbound values to the second stage goroutine
func randomCounter(out chan<- int) {
	defer wg.Done()
	var random int
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random // blocks until the receiving channel reacts, then the loop continues
	}
	close(out)
}

// Produces Fibonacci values of inputs provided by randomCounter
// Second stage
func generateFibonaci(out chan<- fibvalue, in <-chan int) {
	defer wg.Done()
	var input float64

	for v := range in {
		input = float64(v)
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
		out <- fibvalue{
			input: v,
			value: int(result),
		}
	}
	close(out)
}

// Third stage of the pipeline
func printFibonacci(in <-chan fibvalue) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Fibonacci value of %d is %d\n", v.input, v.value)
	}
}

func main() {
	// Add 3 into WaitGroup Counter
	wg.Add(3)
	// Declare channels (unbuffered)
	randoms := make(chan int)
	fibs := make(chan fibvalue)
	// Launching 3 goroutines
	go randomCounter(randoms) // First stage of pipeline
	go generateFibonaci(fibs, randoms) // Second stage of pipeline
	go printFibonacci(fibs) // Third stage of pipeline
	// Wait of completing all goroutines
	wg.Wait()
}
