package main

import (
		"fmt"
)

//The defer keyword in Go allows us to defer the execution of a 
//function call until the return of the parent function.
func main() {
	defer goodbye()
	fmt.Println("hello")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

//Deferred calls will still get 
//called even if another deferred call panics.
	defer fmt.Println("one")
	defer panic("two")
	defer fmt.Println("three")
}

func goodbye() {
	fmt.Println("goodbye")
}

//Last in first out LIFO 
//output:
//hello
//goodbye
//three
//two
//one

//three
//one
//panic: two