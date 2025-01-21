package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// defer fmt.Println("Bye")
	// defer fmt.Println("Yellow")
	// defer fmt.Println("black")
	// fmt.Println("Hello")
	// fmt.Println("Nice")
	f, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	val, _ := io.ReadAll(f)
	fmt.Println(string(val))
}
