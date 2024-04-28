// to run files

// go run main.go deck.go oddOrEven.go state.go

// to test file

// go test

package main

import "fmt"

func main() {

	fmt.Println("Hello!")

	// // cards := deck{"six of spade", "Jack of Diamonds"}

	// cards.saveToFile("tmp")

	// // cards2 := newDeckFromFile("my")
	// // Error :  open my: no such file or directory
	// // exit status 1

	// cards2 := newDeckFromFile("tmp")

	// cards2.print()

	cards3 := newDeck()

	cards3.suffle()

	cards3.print()

	// below lines are for utilising some functions

	cards := newDeck()

	fmt.Println(cards.toString())

	hand, ramainingCard := deal(cards, 5)

	hand.print()

	ramainingCard.print()

	cards.print()

	// basics
	printOddOrEven()

	printState()

	arr()

}
