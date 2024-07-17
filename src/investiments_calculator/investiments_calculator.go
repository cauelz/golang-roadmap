package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.0

func main() {

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

	futureValue := calculateFutureValue(investmentAmount, interestRate, years)

	futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)

	fmt.Print("Future real value is ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, interestRate, years float64) float64 {
	return investmentAmount * math.Pow((1+interestRate/100), years)
}

func calculateFutureRealValue(futureValue, inflationRate, years float64) float64 {
	return futureValue / math.Pow((1+inflationRate/100), years)
}
