package main

import (
	"fmt"
)

func main() {
	var card string = "Ace of Card" // this is way one to declare variable

	var card1 = "Ace of Card1" // way two

	name := "Manash Mondal" // way three

	fmt.Println(card)
	fmt.Println(card1)
	fmt.Println(name)

	card = "value change card1"
	card1 = "value change card2"
	name = "Kumar Manash"

	fmt.Println(card)
	fmt.Println(card1)
	fmt.Println(name)

}
