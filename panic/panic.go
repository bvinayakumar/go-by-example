package main

import "fmt"

func main() {
	// A panic typically means something went unexpectedly wrong.
	// Mostly we use it to fail fast on errors that shouldn’t occur during
	// normal operation, or that we aren’t prepared to handle gracefully.
	panic("a problem")

	fmt.Println("unreachable code")
}
