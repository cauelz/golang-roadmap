package structbasics

import (
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {

	// Create a new person
	// Is also possible to create a person in the following way:
	// p := person{"John", "Doe"}
	p := person{firstName: "John", lastName: "Doe"}

	fmt.Println(p)

	if p.firstName != "John" {
		t.Errorf("Expected John, got %v", p.firstName)
	}
	if p.lastName != "Doe" {
		t.Errorf("Expected Doe, got %v", p.lastName)
	}
}

func TestNewPersonZeroValues(t *testing.T) {

	// Create a new person with zero values
	var p person

	// What are zero values?
	// The zero value is the default value of a type.
	// For example, the zero value of an int is 0, and the zero value of a string is an empty string.
	// For booleans, the zero value is false.
	// The zero value of a struct is a struct with all its fields set to their zero values.
	// The zero value of a pointer is nil.
	fmt.Println(p)

	// Print the zero values with the fmt package and the Printf function
	fmt.Printf("First Name: %v\n", p.firstName)
	fmt.Printf("Last Name: %v\n", p.lastName)
	fmt.Printf("Age: %v\n", p.age)
	fmt.Printf("Is Adult: %v\n", p.isAdult)


	if p.firstName != "" {
		t.Errorf("Expected empty string, got %v", p.firstName)
	}
	if p.lastName != "" {
		t.Errorf("Expected empty string, got %v", p.lastName)
	}
}


