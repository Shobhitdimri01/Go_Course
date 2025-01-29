package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	// Simulate some work
	time.Sleep(2 * time.Second)

	// Send data to the channel
	ch <- "Task completed"

}

func main() {
	// Create a channel of type string
	ch := make(chan string)

	// Start a goroutine that sends data into the channel
	go worker(ch)

	// Receive data from the channel (this blocks until data is received)
	fmt.Println("Waiting for the worker to finish...")
	message := <-ch
	fmt.Println("Received:", message)
}
