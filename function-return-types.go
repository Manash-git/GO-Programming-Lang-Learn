package main

import "fmt"

func main() {

	fmt.Println("First GO function test program")

	card := newCard() // return value is assigned to card
	fmt.Println(card) // printing the value of card

	cardNo := number()
	fmt.Println(cardNo)

}

func newCard() string { // types string needed to be declared
	return "Five of Diamonds"
}

func number() int {
	return 420
}
