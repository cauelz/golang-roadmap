package main

import "testing"


func TestEvenOrOdd(t *testing.T) {
	result := evenOrOdd(5)

	if result != "Odd" {
		t.Errorf("Expected Odd, got %v", result)
	}

	result = evenOrOdd(4)

	if result != "Even" {
		t.Errorf("Expected Even, got %v", result)
	}
}