package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	SimpleFlags()
	// Create a variable of type IntSlice
	// var numbers IntSlice

	// // Register the flag with the flag package
	// flag.Var(&numbers, "numbers", "Comma-separated list of integers")

	// // Parse the command-line flags
	// flag.Parse()

	// // Output the parsed flag value
	// fmt.Println("Parsed numbers:", numbers)

	// //go run main.go --numbers 10,20,30,40

	// //go run main.go --set key1=value1 --set key2=value2

	// var settings KeyValue = make(map[string]string)

	// flag.Var(&settings, "set", "Key-value pair in the format key=value")
	// flag.Parse()

	// fmt.Println("Parsed settings:", settings)

}

// IntSlice is a custom type that implements the flag.Value interface
type IntSlice []int

// String returns the string representation of the value (used for displaying the flag value)
func (i *IntSlice) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set parses and sets the value from a string
func (i *IntSlice) Set(value string) error {
	// Split the input string by commas
	parts := strings.Split(value, ",")
	for _, part := range parts {
		// Convert each part to an integer
		num, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("invalid number: %s", part)
		}
		*i = append(*i, num)
	}
	return nil
}

//key value

type KeyValue map[string]string

func (kv *KeyValue) String() string {
	return fmt.Sprintf("%v", *kv)
}

func (kv *KeyValue) Set(value string) error {
	parts := strings.Split(value, "=")
	if len(parts) != 2 {
		return fmt.Errorf("invalid key-value pair: %s", value)
	}
	key, val := parts[0], parts[1]
	(*kv)[key] = val
	return nil
}
