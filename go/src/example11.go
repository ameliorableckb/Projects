package main

import (
		"fmt"
)


// Create a function called `printSlice` that takes a slice of strings and prints them out.
// func () {}
func printSlice(sl []string) {
	for _, s := range sl {
		fmt.Println(s)
	}
}

// Create a function called `printAll` that takes a string as a variadic argument and prints them out.
// func () {}
func printAll(sl ...string) {
	for _, s := range sl {
		fmt.Println(s)
	}
}

func main() {
	fruit := []string{"banana", "orange", "strawberry", "apple"}
	fmt.Println("Print out a slice")
	printSlice(fruit)
	println() // insert blank line in output

	fmt.Println("Print out multiple items")
	printAll("banana", "orange", "strawberry", "apple")

	println() // insert blank line in output
	fmt.Println("Print out multiple items")
	printAll("banana", "orange", "strawberry", "apple")

	// Explode the arguments to pass a slice to printAll
	println() // insert blank line in output
	printAll(fruit...)

	// You can also call printAll with no arguments
	println() // insert blank line in output
	printAll()
}
