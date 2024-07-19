package struct_basics

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewPerson(t *testing.T) {

	// Create a new person
	// Is also possible to create a person in the following way:
	// p := person{"John", "Doe"}
	p := person{firstName: "John", lastName: "Doe"}

	fmt.Println("--- Person creation ---")
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
	fmt.Println("--- Person zero values ---")
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

func TestNewPersonWithContact(t *testing.T) {

	// Create a new person with contact info
	p := person{
		firstName: "John",
		lastName: "Doe",
		contactInfo: contact{
			email: "john.doe@example.com",
			phone: "123-456-7890",
			zipCode: 12345,
		},
	}

	fmt.Println("--- Person with contact info ---")
	fmt.Println(p)

	// checks if person has valid email
	if !strings.Contains(p.contactInfo.email, "@") {
		t.Errorf("Expected valid email, got %v", p.contactInfo.email)
	}

	// checks if person has a non empty phone number
	if len(p.contactInfo.phone) == 0 {
		t.Errorf("Expected filled phone number, got %v", p.contactInfo.phone)
	}

	// checks if person has a non empty zip code
	if p.contactInfo.zipCode == 0 {
		t.Errorf("Expected filled zip code, got %v", p.contactInfo.zipCode)
	}
}

func TestPrint(t *testing.T) {

	p := person{firstName: "John", lastName: "Doe"}

	fmt.Println("--- Print person ---")
	p.print()
}

func TestUpdateName(t *testing.T) {

	p := person{firstName: "John", lastName: "Doe"}

	fmt.Println("--- Update name ---")
	fmt.Println("Before update:")
	fmt.Println(p)

	p.updateName("Jane")

	fmt.Println("After update:")
	fmt.Println(p)

	if p.firstName != "Jane" {
		t.Errorf("Expected Jane, got %v", p.firstName)
	}
}

func TestGettinUserInputAndCreatePerson(t *testing.T) {

}