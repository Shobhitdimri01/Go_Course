package main

/*
#include <math.h>
*/
import "C"
import "fmt"

func main() {
	result := C.sqrt(16.0) // Call sqrt from math.h
	fmt.Println("Square root of 16 is:", result)
}
