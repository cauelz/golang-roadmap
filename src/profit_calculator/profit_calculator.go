package main

import "fmt"



func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter revenue: ")

	expenses = getUserInput("Enter expenses: ")

	taxRate = getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// fmt.Println("Earnings before taxes: ", ebt)
	// fmt.Println("Profit: ", profit)
	// fmt.Println("Profit ratio: ", ratio)

	//fmt.Printf("Earnings before taxes: %.2f\n", ebt)
	//fmt.Printf("Profit: %.2f\n", profit)
	//fmt.Printf("Profit ratio: %.2f\n", ratio)

	ebtString := fmt.Sprintf("Earnings before taxes: %.2f\n", ebt)
	profitString := fmt.Sprintf("Profit: %.2f\n", profit)
	ratioString := fmt.Sprintf("Profit ratio: %.2f\n", ratio)

	fmt.Println("ebtString: ", ebtString)
	fmt.Println("profitString: ", profitString)
	fmt.Println("ratioString: ", ratioString)

}

func getUserInput(text string) (input float64) {
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

func calculateEBT(revenue, expenses float64) (ebt float64) {
	ebt = revenue - expenses
	return ebt
}

func calculateProfit(ebt, taxRate float64) (profit float64) {
	profit = ebt * (1 - taxRate/100)
	return profit
}

func calculateProfitRatio(profit, revenue float64) (ratio float64) {
	ratio = profit / revenue
	return ratio
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = calculateEBT(revenue, expenses)
	profit = calculateProfit(ebt, taxRate)
	ratio = calculateProfitRatio(profit, revenue)
	return ebt, profit, ratio
}