package pointer

import (
	"testing"
)


func TestPassByValues(t *testing.T) {

	age := 40

	adultAge := getAdultYearsByValue(age)

	if age != 40 {
		t.Errorf("Expected 40 but got %d", age)
	}

	if(adultAge != 22) {
		t.Errorf("Expected 22 but got %d", adultAge)
	}
}

func TestPassByPointers(t *testing.T) {

	age := 40

	adultAge := getAdultYearsByPointer(&age)

	if age != 22 {
		t.Errorf("Expected 22 but got %d", age)
	}

	if(adultAge != 22) {
		t.Errorf("Expected 22 but got %d", adultAge)
	}
}