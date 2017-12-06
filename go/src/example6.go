package main

import "fmt"

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	fmt.Println("1")
	// Loop through the map and print out the key and the value
	for key, value := range beatles {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	fmt.Println("2")
	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	bob, ok := beatles["Bob"]
	if !ok {
		fmt.Println("not found") 
	}else {
		fmt.Println(bob)
	}

}