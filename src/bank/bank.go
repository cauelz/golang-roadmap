package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Account struct {
	AccountNumber int
	AccountBalance float64
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 0, errors.New("Error reading file!")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Error parsing float!")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	fmt.Println("Writing Account Balance to File...")
	balanceText := fmt.Sprint(balance)
	// What is this file permission?
	// 0644 is a file permission in Unix-like systems. It is a 3-digit octal number.
	// The first digit is the owner's permission, the second digit is the group's permission, and the third digit is the others' permission.
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error getting balance from file!")
		fmt.Println(err)
		fmt.Println("---------")
		// When to use panic?
		// panic is used to terminate the program immediately. It is used when a function returns an error that the program cannot recover from.
		//panic(err)
	}	
	
	fmt.Println("Welcome to the Go Bank!")
	// Conditional for loops in Go can be written as follows:
	// for condition {
	// 	// Code to be executed
	// }
	// Where condition is a boolean expression
	// Infinite loop
	for {
	
		fmt.Println("Please select an option:")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Check Balance")
		fmt.Println("5. Exit")
	
	
		var option int
		fmt.Scanln(&option)
	
		switch option {

		case 1:
			fmt.Println("Enter Account Number:")
			var accountNumber int
			fmt.Scanln(&accountNumber)
			fmt.Println("Enter Initial Deposit:")
			var initialDeposit float64
			fmt.Scanln(&initialDeposit)
			accountBalance = initialDeposit
			fmt.Println("Account Created Successfully!")
		case 2:
			fmt.Println("Enter Deposit Amount:")
			var depositAmount float64
			fmt.Scanln(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Deposit Amount should be greater than 0!")
				continue
			}
	
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Deposit Successful!")
		case 3:
			fmt.Println("Enter Withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scanln(&withdrawalAmount)
	
			if withdrawalAmount <= 0 {
				fmt.Println("Withdrawal Amount should be greater than 0!")
				continue
			}
	
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient Funds!")
				fmt.Println("Your Account Balance is: ", accountBalance)
				continue
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Withdrawal Successful!")
			}
		case 4:
			fmt.Println("Account Balance: ", accountBalance)
		case 5:
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid Option!")
		}
	
	}

}