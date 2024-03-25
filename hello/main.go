package main

import (
	"fmt"
)


func main() {
		// Call the function
		words := []string{"hello", "world"}
		hello(words)
}

// Function definition
func hello(words []string) {
		// Print the message

		for _, word := range words {
			fmt.Println(word)
		}
}