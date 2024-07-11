package structbasics

// What is a Struct in Golang?
// A struct is a user-defined type that contains a collection of named fields/properties.
// It is used to group related data together to form a single unit.
// It is similar to a class in object-oriented programming languages, but it does not have inheritance.

type person struct {

	firstName string
	lastName string
	age int
	isAdult bool
	contactInfo contact	// Embedded struct
}

type contact struct {

	email string
	phone string
	zipCode int
}

