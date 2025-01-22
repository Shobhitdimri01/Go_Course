package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)
	printValue(5)
	printValue("Golang")
	printValue([]int{1, 2, 3, 4})
}
func printValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Square of", v, "is", v*v)
	case string:
		fmt.Println("String length of", v, "is", len(v))
	case []int:
		fmt.Println("Sum of slice:", sumSlice(v))
	default:
		fmt.Println("Unsupported type:", v)
	}
}

func sumSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}
