package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)  //Ringo
}

func printValue(s *string) {
	// print the pointer value

	fmt.Println(s)
	// print the string value
	fmt.Println(*s) //George
	// change the string value
	*s = "Ringo"
}