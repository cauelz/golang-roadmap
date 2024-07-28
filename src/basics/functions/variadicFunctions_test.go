package functions

import (
	"fmt"
	"testing"
)


func TestSumValues(t *testing.T) {

	summedValue := SumValues(1, 2, 3, 4, 5)

	fmt.Println("TestSumValues 1: ", summedValue)

	values := []int{1, 2, 3, 4, 5}

	summedValue = SumValues(1, values...)

	fmt.Println("TestSumValues 2: ", summedValue)
}