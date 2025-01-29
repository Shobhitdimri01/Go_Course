package main

import "fmt"

func main() {
	var a int
	var p string
	fmt.Println("Please enter a number & letter:")
	fmt.Scanf("%d %s", &a, &p)
	fmt.Println("My value:", a, "\nMy string:", p)
}
