package main

import "fmt"

func main() {
	greet := func(name string) string {
		return "Hello " + name
	}

	msg := greet("Raj")
	fmt.Println(msg)

	increment := func(x int) int {
		return x + 1
	}
	fmt.Println(increment(2))
}
