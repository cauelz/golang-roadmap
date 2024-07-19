package main

import (
	"errors"
	"fmt"
	"os"
)



func main() {

    revenue, revenueError := getUserInput("Enter revenue: ")

    expenses, expensesError := getUserInput("Enter expenses: ")

    taxRate, taxrateError := getUserInput("Enter tax rate: ")

    if revenueError != nil || expensesError != nil || taxrateError != nil {
        fmt.Println("Error ", revenueError, expensesError, taxrateError)
		// panic(revenueError) can be used to terminate the program immediately and print the error message
        return
    }

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

	err := saveToFile("financials.txt", ebtString+profitString+ratioString)

	if err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}

}

func saveToFile(filename, data string) (err error) {

	err = os.WriteFile(filename, []byte(data), 0644)

	if err != nil {
		return err
	}
	return nil
}

func getUserInput(text string) (input float64, err error) {

	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		err = errors.New("invalid input! please enter a positive number")

		return 0, err
	}

	return input, nil
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