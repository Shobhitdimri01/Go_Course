package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	s := []string{"apple", "dkc", "dhjs", "sksk", "skskd", "kfkf", "sksks"}
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go printMydata(s, &wg)
	go PrintAgain(s, &wg)
	wg.Wait()
	fmt.Println(time.Since(t))
}

//1 thread -> sequential programming
//2 thread -> Concurrent programming -> t1   t2
// Parallel programming

func printMydata(s []string, wg *sync.WaitGroup) {
	for i, v := range s {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d - %s\n", i, v)
	}
	wg.Done()
}

func PrintAgain(s []string, wg *sync.WaitGroup) {
	for i, v := range s {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d - %s\n", i, v)
	}
	wg.Done()
}
