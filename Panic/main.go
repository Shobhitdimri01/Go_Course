package main

import "fmt"

func main() {
	fmt.Println("hello")
	var p *int
	*p=10
	fmt.Println(p)
}
