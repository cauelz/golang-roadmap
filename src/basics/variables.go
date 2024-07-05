package main

import "fmt"

// Go is a statically typed language, which means that we need to declare the type of a variable
// when we declare it. We can use the following syntax to declare a variable:
// var <variable_name> <type> = <value>

// We can also declare variables outsite of the main function, but we need to use the var keyword
// and we can't use the := syntax
var deckSize int

func variables() {

	// deckSize can be initialized here, but it can also be initialized outside of the main function
	deckSize = 52

	// var card string = "Ace of Spades" this is the long way to declare a variable
	// in Go, we can use the following shorthand to declare a variable and assign a value to it
	// this is called type inference and Go will automatically infer the type of the variable
	card := "Ace of Spades"

	// We can also reassign the value of a variable
	card = "Five of Diamonds"

	// We can also declare a variable without assigning a value to it
	var card2 string
	card2 = "Two of Hearts"

	// We can also declare multiple variables at once
	// var card, card2 string = "Ace of Spades", "Two of Hearts"
	// card, card2 := "Ace of Spades", "Two of Hearts"
	

	fmt.Println(card)
	fmt.Println(card2)
}