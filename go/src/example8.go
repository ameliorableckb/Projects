package main

import (
		"fmt"
)

func main() {
	sayHello2("Hello", "Mark")
	fmt.Println(sayHello())

	l, c := slicer([]string{"John", "Paul", "George", "Ringo"})
	fmt.Println(l)
	fmt.Println(c)

	d := exists()
	fmt.Println(d)

}

//Functions can return 0 to N numbers of values, 
//though it is considered best practice to not
// return more than two or three.
func sayHello2(greeting string, name string) {
	fmt.Printf("%s, %s", greeting, name)
}

//The type of the return value(s) is declared after the function definition.
func sayHello() string {
	return "Hello"
}

func slicer(s []string) (int, int) {
	return len(s), cap(s)
}

//Named Returns
//It is best practice to not use named returns, 
//unless they are needed for documentation.
//Multiple same types etc.
func exists() (exist bool) {
    exist = true
    return
}