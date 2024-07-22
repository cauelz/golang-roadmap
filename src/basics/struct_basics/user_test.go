package struct_basics

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {

	var user *User = New("John", "Doe", "john.doe@example.com", "password", "25")


	if user.firstName != "John" {
		t.Errorf("Expected %v, got %v", "John", user.firstName)
	}

	if user.lastName != "Doe" {
		t.Errorf("Expected %v, got %v", "Doe", user.lastName)
	}

	if user.email != "john.doe@example.com" {
		t.Errorf("Expected %v, got %v", "john.doe@example.com", user.email)
	}

	if user.password != "password" {
		t.Errorf("Expected %v, got %v", "password", user.password)
	}
	
	if user.age != "25" {
		t.Errorf("Expected %v, got %v", 25, user.age)
	}
}

func TestOutputUserDetails(t *testing.T) {
	var user *User = New("John", "Doe", "john.doe@example.com", "password", "25")

	OutputUserDetails(user)

}

func TestDeferrencingPointers(t *testing.T) {
	var user *User = New("John", "Doe", "john.doe@example.com", "password", "25")

	fmt.Println("First Name: ", (*user).firstName)
	fmt.Println("Last Name: ", user.lastName)

}

func TestNewAdminUser(t *testing.T) {

	var user *Admin = NewAdmin("John", "Doe", "john.doe@example.com", "password", "25", "admin")

	if user.firstName != "John" {
		t.Errorf("Expected %v, got %v", "John", user.firstName)
	}

	if user.lastName != "Doe" {
		t.Errorf("Expected %v, got %v", "Doe", user.lastName)
	}

	if user.email != "john.doe@example.com" {
		t.Errorf("Expected %v, got %v", "john.doe@example.com", user.email)
	}

	if user.password != "password" {
		t.Errorf("Expected %v, got %v", "password", user.password)
	}

	if user.age != "25" {
		t.Errorf("Expected %v, got %v", 25, user.age)
	}

	if user.role != "admin" {
		t.Errorf("Expected %v, got %v", "admin", user.role)
	}
}

	
