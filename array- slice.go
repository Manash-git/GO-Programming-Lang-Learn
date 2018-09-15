package main

import (
	"fmt"
)

func main() {
	cards := []string{newCard(), "Hello World"}
	fmt.Println("Direct Print => ", cards)

	//? Append() is used for modified the existing array/slice and return a new array

	cards = append(cards, "|| Append use here ||", newCard(), "|| Array list ")
	fmt.Println(cards)

	name := []string{"i", "am", "manash"}
	fmt.Println(name)

	/*
		* Below line isn't valid
		! name = append(name, cards, "Hello World")
	*/
}

func newCard() string {
	return "Hi there"
}
