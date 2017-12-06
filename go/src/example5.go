package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	// Create a new slice called family by joining the parents and kids slice together
	//dots are needed to append a second slice
	//https://blog.golang.org/slices
	family := append(parents, kids...)

	fmt.Println(family)

	// Fix the following bugs
	//slices can be made with the make command, starting length and capacity
	
	//extras := make([]string, 1, 0)
	extras := make([]string, 1, 1)
	//	extras[1] = "Alice"
	extras[0] = "Alice"
	fmt.Println(extras)
}