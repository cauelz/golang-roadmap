package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("\n\n\nHand of cards:")
	hand.print()

	fmt.Println("\n\n\nRemaining deck:")
	remainingDeck.print()

	fmt.Println("\n\n\nDeck as a string:")
	fmt.Println(cards.toString())

	fmt.Println("\n\n\nSaving deck to file...")
	cards.saveToFile("my_deck")
}
