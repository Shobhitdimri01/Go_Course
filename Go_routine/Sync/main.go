package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	// fmt.Println("cores:", runtime.NumCPU())
	// runtime.GOMAXPROCS(2)

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers completed")
	fmt.Println("total time taken ", time.Since(t))
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	fmt.Printf("Worker %d done\n", id)
}
