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



func main() {
	fmt.Println("1.")
	//Functions As Variables
	f := func() {
		fmt.Println("Hello")
	}

	f() // Hello

	fmt.Println("2.")
//function has access to arguments declared before it's declaration.
	name := "Ringo"
	g := func() {
		fmt.Printf("Hello %s\n", name)
	}

	g() // Hello Ringo

	fmt.Println("3.")
//Functions in Go can be passed as values to other functions.
	h := func() string {
		return "Hello"
	}
	sayHello(h)

	fmt.Println("4.")
//Anonymous functions are functions that aren't assigned 
//to a variable, or that don't have a "name" associated with them.
	sayHello2(func() string {
		return "Hello"
	})

}

func sayHello(fn func() string) {
	fmt.Println(fn())
}

type greeter func() string

func sayHello2(fn greeter) {
	fmt.Println(fn())
}