package struct_basics

import "fmt"

// What is a Struct in Golang?
// A struct is a user-defined type that contains a collection of named fields/properties.
// It is used to group related data together to form a single unit.
// It is similar to a class in object-oriented programming languages, but it does not have inheritance.

type person struct {

	firstName string
	lastName string
	age int
	isAdult bool
	// It is possible to embed a struct information in another struct with the following syntax:
	// contact	// Embedded struct without a field name
	contactInfo contact	// Embedded struct
}

type contact struct {
	email string
	phone string
	zipCode int
}

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scan(&input)
	return input
}