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
	
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter interest rate: ")
	fmt.Scan(&interestRate)

	outputText("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+interestRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Print("Future real value is ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}
