package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	cmdLineOutput := os.Args
	a := cmdLineOutput[1]
	b := cmdLineOutput[2]
	fmt.Println(Add(a,b))

}

func Add(a, b string) int {
	m, _ := strconv.Atoi(a)
	n, _ := strconv.Atoi(b)
	return m + n

}
