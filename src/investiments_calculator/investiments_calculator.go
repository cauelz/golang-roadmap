package main

import (
	"fmt"
	"math"
)

func main() {

	// How to declare Variables and Constants in Go
	// var <variable_name> <variable_type> = <value> declare and assign the value to the variable
	// const <constant_name> <constant_type> = <value> declare and assign the value to the constant
	// var <variable_name> <variable_type> just declare the variable
	// <variable_name> = <value> assign the value to the variable
	// <variable_name> := <value> declare and assign the value to the variable (only works inside functions) and Go will infer the type of the variable
	const inflationRate = 3.0
	var investmentAmount float64 = 1000
	interestRate, years  := 5.5, 5.0

	futureValue := investmentAmount * math.Pow((1 + interestRate/100), years)

	futureRealValue := futureValue / math.Pow((1 + inflationRate/100), years)

	fmt.Print("Future real value is ", futureRealValue)
}