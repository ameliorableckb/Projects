package main

import (
		"fmt"
		"time"
)

// Create a function called `echo` that takes a 
//single argument of type string and prints it out.
// func () {}
func echo(s string) {
	fmt.Println(s)
}

// Convert this function to use a named return 
//and return it as the default.
func firstDayOfWeek(t time.Time) bool {
	firstDay := t.Weekday() == 1
	return firstDay
}

func main() {
	now := time.Now()
	fmt.Printf("today is the first day of the week: %v", firstDayOfWeek(now))
}
