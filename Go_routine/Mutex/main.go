package main

import (
	"sync"
)

var counter int

// var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	//mu.Lock()   // Lock the mutex
	counter++ // Critical section
	//mu.Unlock() // Unlock the mutex
}

func main() {
	// var wg sync.WaitGroup

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go increment(&wg)
	// }

	// wg.Wait()
	// fmt.Println("Final Counter:", counter)
	RaceExample()
}
