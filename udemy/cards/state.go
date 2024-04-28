package main

import "fmt"

// tmp := 30

// var tmp int

func printState() {

	// // var tmp int
	// tmp = 30

	// fmt.Println(tmp)

	// card  := "Ace of Spades" // only use for initialisation not for re-assining

	var card string = "Ace of Spades"

	fmt.Println(card)

	card = newCard()

	fmt.Println(card)

	fmt.Println("Bangaluru")
}

func newCard() string {
	return "Five of Diamonds"
}

// Array - fixed length list of things
// Slices - An array that can grow or shrink

func arr() {
	cards := []string{newCard(), "Jack of Diamonds"}

	cards = append(cards, "six of spade")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}
