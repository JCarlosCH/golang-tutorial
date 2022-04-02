package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Juan")
	fmt.Println(message)
}
