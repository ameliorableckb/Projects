package main

import (
		"fmt"
		"os"
	//  "log"
)

//The defer keyword in Go allows us to defer the execution of a 
//function call until the return of the parent function.
func main() {
	defer fmt.Println("one")
	os.Exit(1)
//	log.Fatal("boom")
	defer fmt.Println("three")
}

func goodbye() {
	fmt.Println("goodbye")
}

//Deferred calls will not fire 
//if the code explicitly calls os.Exit or log.Fatal

//There will be no output of one or three