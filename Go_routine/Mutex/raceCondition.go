package main

import (
	"fmt"
	"sync"
)

// to check race condition : go run -race *

var msg string

func UpdateMsg(s string, wg *sync.WaitGroup,m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg += " " + s
	m.Unlock()
}

func RaceExample() {
	var wg sync.WaitGroup
	var m sync.Mutex
	msg = "Hello, World"
	wg.Add(2)
	go UpdateMsg("Hello Cosmos", &wg,&m)
	go UpdateMsg("Hello Universe", &wg,&m)
	wg.Wait()

	fmt.Println(msg)
}
