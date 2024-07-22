package struct_basics

import (
	"errors"
	"fmt"
)

// To keep the fields of the struct private, we can use the lowercase letter
// To keep the fields of the struct public, we can use the uppercase letter
type User struct {
	// Define the fields of the struct, which are the properties of the struct
	firstName string
	lastName string
	age string
	email string
	password string
}

// Define a struct that inherits from the User struct
// The Admin struct has all the fields of the User struct and an additional field called role
// The User methods can be accessed by the Admin struct, example, the ClearUserDetails method
type Admin struct {
	// The Admin struct inherits the fields of the User struct and defining without an alias allows us to acces the fields and methods of the
	// User struct directly
	User
	role string
}

func (u *User) ClearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.age = ""
	u.email = ""
	u.password = ""
}

/*
* @description : This function creates a new user of type Admin.
* @param : firstName, lastName, email, password, age, role
* @return : *Admin
*/
func NewAdmin(firstName, lastName, email, password, age, role string) (*Admin) {
	
	// Validate the user input
	err := validateUserInput(firstName, lastName, email, password, age)

	if err != nil {
		fmt.Println(err)
		return nil // Return nil if the user input is invalid
	}

	var a Admin = Admin{
		User: User{
			firstName: firstName,
			lastName: lastName,
			age: age,
			email: email,
			password: password,
		},
		role: role,
	}

	return &a
}

// The struct fields are private, so we need to define methods to access the fields and populate the fields
// The methods are defined outside the struct, but they are associated with the struct
func New(firstName, lastName, email, password, age string) (*User) {

	// Validate the user input
	err := validateUserInput(firstName, lastName, email, password, age)

	if err != nil {
		fmt.Println(err)
		return nil // Return nil if the user input is invalid
	}

	var u User = User{}
	u.firstName = firstName
	u.lastName = lastName
	u.age = age
	u.email = email
	u.password = password

	return &u
}

func validateUserInput(firstName, lastName, email, password, age string) error {

	if firstName == "" {
		return errors.New("First name is required")
	}

	if lastName == "" {
		return errors.New("Last name is required")
	}

	if email == "" {
		return errors.New("Email is required")
	}

	if password == "" {
		return errors.New("Password is required")
	}

	if age == "" {
		return errors.New("Age is required")
	}

	return nil
}

func OutputUserDetails(u *User) {
	// It is also possible to access a pointer value using the dot operator
	// This is because Go automatically dereferences the pointer

	// Output the user details
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Age: ", u.age)
	fmt.Println("Email: ", u.email)
	fmt.Println("Password: ", u.password)

	//this also works
	//fmt.Println("First Name: ", (*u).firstName)
}

func GetInfoFromUserInput(promptMessage string) (string) {
	
	var input string

	fmt.Println(promptMessage)
	// Read the user input
	//fmt.Scan(&input)

	// Read the user input until the user presses the enter key
	fmt.Scanln(&input)

	return input
}

