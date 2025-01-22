package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This is an error : ")
		}
	}()
	// fmt.Println("hello")
	// var p *int
	// *p = 10
	// fmt.Println(p)
	s := []int{2, 3, 5, 7}
	for i := 0; i < 10; i++ {
		fmt.Println(s[i])
	}
}
