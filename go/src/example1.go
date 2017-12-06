package main

import "fmt"

func main() {
	// Declare the following variables with zero values:
	// not zero but no value
	// name: command, type string
	// name: valid, type bool
	//command := "0"
	//valid := 0
	var command string
	var valid bool


	// print the value of those variables
    fmt.Printf("command = %s\n", command)
    fmt.Printf("valid = %v\n", valid)

	// declare and initialize the following variables:
	// name: foo, type: int, value: 5
	// name bar, type: bool, value: true
    foo := 5
    bar := true

	// print the value of those variables
    fmt.Printf("foo = %d\n", foo)
    fmt.Printf("bar = %v\n", bar)

}