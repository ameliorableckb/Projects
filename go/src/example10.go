package main

import (
		"fmt"
)


func main() {
	sayHello("John", "Paul", "George", "Ringo")
	// Hello John
	// Hello Paul
	// Hello George
	// Hello Ringo
	sayHello("George")
	// George
	sayHello()
	//
	beatles := []string{"John", "Paul", "George", "Ringo"}
	sayHello2(beatles...)
}

//Variadic arguments must be the last 
//argument of the function.
func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}

func sayHello2(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}