package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 3.0
	//var investmentAmount, interestRate, years float64
	var (
		investmentAmount float64
		interestRate     float64
		years            float64
	)
	
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter interest rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+interestRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Print("Future real value is ", futureRealValue)
}
