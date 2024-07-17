package main

import "fmt"



func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / revenue

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

