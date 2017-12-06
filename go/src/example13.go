package main

import (
		"fmt"
)


type Beatle struct {
	Name       string
	Instrument string
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Drums"}
	fmt.Println(b.String()) // Ringo plays Drums
}

func (b Beatle) String() string {
	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
}

//There isnt method overloading. You can only have one string()
//but if you attach it to reciever such as (b Beatle)
//You can have multiple String() per item
//They operate the same

//func (b Beatle) String() string {
//	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
//}

//func String(b Beatle) string {
//	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
//}