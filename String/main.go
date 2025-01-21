package main

import (
	"fmt"
	"strings"
)

// ASCII
func main() {
	// a := 'h'
	// var p string
	// l := "hello"
	// p = "red\tyellow"
	// fmt.Println(string(a), "\nstring:", p)
	// is := strings.ContainsRune(l, a)
	// fmt.Printf("does %s contain h ? %v\n", l, is)
	// Check()
	//StringSlice()
	Compare()
}

func Check() {
	p := "abcdef"

	for i, j := range p {
		fmt.Printf("%d - %v\n", i, string(j))
	}
}

func StringSlice() {
	s := []string{"a", "b", "c", "d", "e"}
	s = append(s[0:2], s[3:]...)
	fmt.Println(s)

}

func Compare() {
	a := "apple"
	b := "AppLe"
	fmt.Println(strings.EqualFold(a, b))
}
