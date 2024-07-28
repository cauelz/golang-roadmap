package functions

import (
	"fmt"
	"testing"
)


func TestFactorial(t *testing.T) {
	fmt.Println("Factorial of 5 is", factorial(5))
}