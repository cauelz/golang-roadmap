package functions

import (
	"fmt"
	"testing"
)


func TestSumValues(t *testing.T) {

	summedValue := SumValues(1, 2, 3, 4, 5)

	fmt.Println("TestSumValues: ", summedValue)
}