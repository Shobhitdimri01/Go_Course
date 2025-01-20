package main

import "fmt"

func main() {
	c1 := complex(2, 3)
	c2 := complex(2, 3)
	c3 := c1+c2
	fmt.Println(c3)
	fmt.Println("real:", real(c1))
	fmt.Println("imaginary", imag(c1))
}
