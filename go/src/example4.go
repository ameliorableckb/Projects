package main

import (
		"fmt"
)

func main() {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.

	//fmt.Printf(fruits)
	fmt.Println("1.")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Fruit: %s, Index: %v \n", fruits[i], i)
	}

	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	fmt.Println("2.")
	for i, n := range fruits {
		fmt.Printf("Fruit: %s, Index: %v \n", n, i)
	}

	// Using the keyword `continue`, skip every other fruit (even ones)
	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	fmt.Println("3.")
	for i, n := range fruits {
		//modulos to get remainder
		if (i%2 == 0 && i != 0){
			continue
		}
		fmt.Printf("Fruit: %s, Index: %v \n", n, i)
	}
}