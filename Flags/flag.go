package main

import (
	"flag"
	"fmt"
)

//go run main.go -port=9090 -debug=true -env=production

func SimpleFlags() {
	// Define flags
	port := flag.Int("port", 8080, "Port to run the server on")
	debug := flag.Bool("debug", false, "Enable debug mode")
	env := flag.String("env", "development", "Application environment (development/production)")

	// Parse flags
	flag.Parse()

	// Use the flag values
	fmt.Printf("Port: %d\n", *port)
	fmt.Printf("Debug mode: %v\n", *debug)
	fmt.Printf("Environment: %s\n", *env)

	// Print remaining arguments
	fmt.Printf("Remaining arguments: %v\n", flag.Args())
}
