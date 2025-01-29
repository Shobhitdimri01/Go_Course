package main

import "fmt"

func printEven(chEven, chOdd chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-chEven // Wait for the signal to print even
		fmt.Println("Even:", i)
		chOdd <- true // Signal the odd goroutine to proceed
	}
}

func printOdd(chEven, chOdd, done chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-chOdd // Wait for the signal to print odd
		fmt.Println("Odd:", i)
		chEven <- true // Signal the even goroutine to proceed
	}
	done <- true
}

func main() {
	chEven := make(chan bool) // Channel to signal the even goroutine
	chOdd := make(chan bool)  // Channel to signal the odd goroutine
	done := make(chan bool)
	// Start goroutines
	go printEven(chEven, chOdd)
	go printOdd(chEven, chOdd, done)

	// Start the sequence with even
	chOdd <- true // Signal the even goroutine to start

	// Block the main goroutine until both goroutines finish
	<-done // Wait for the last even number to print
}
