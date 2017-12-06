package main

// declare a struct called Movie
// add the following fields:
// - Title (string)
// - Released (bool)
// - Length (int)

import (
		"fmt"
)

type Movie struct {
	Title string
	Released bool
	Length int
}

func main() {
	// declare a variable called "movie" of type "Movie"
	var movie Movie

	// Set the Title to "Wizard of Oz"
	movie.Title = "Wizard of Oz"
	// Set the Released variable to "true"
	movie.Released = true
	// Set Length to 125
	movie.Length = 125

	// Print the value of "movie" out
	// hint: you can use fmt.Println(movie)
	fmt.Println(movie)

}