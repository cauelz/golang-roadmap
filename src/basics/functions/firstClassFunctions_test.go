package functions

import "testing"

func TestDoubleNumbers(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := []int{2, 4, 6}

	transformedNumbers := TransformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	for i, number := range transformedNumbers {
		if number != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], number)
		}
	}
}

func TestTripleNumbers(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := []int{3, 6, 9}

	transformedNumbers := TransformNumbers(&numbers, func(number int) int {
		return number * 3
	})

	for i, number := range transformedNumbers {
		if number != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], number)
		}
	}
}