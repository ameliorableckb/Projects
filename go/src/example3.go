package main

import (
		"fmt"
        "foo"
)

func main() {
	var user foo.User{
		user.First = "Connie",
		user.Last = "Bergquist"
	}

	fmt.Println(user)

}