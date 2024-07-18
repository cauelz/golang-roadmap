package main

import "fmt"

type Account struct {
	AccountNumber int
	AccountBalance float64
}

func main() {

	var accountBalance float64 = 0.0
	
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
	
		if option == 1 {
			fmt.Println("Enter Account Number:")
			var accountNumber int
			fmt.Scanln(&accountNumber)
			fmt.Println("Enter Initial Deposit:")
			var initialDeposit float64
			fmt.Scanln(&initialDeposit)
			accountBalance = initialDeposit
			fmt.Println("Account Created Successfully!")
		} else if option == 2 {
			fmt.Println("Enter Deposit Amount:")
			var depositAmount float64
			fmt.Scanln(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Deposit Amount should be greater than 0!")
				return
			}
	
			accountBalance += depositAmount
			fmt.Println("Deposit Successful!")
			
		} else if option == 3 {
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
			
		} else if option == 4 {
			fmt.Println("Account Balance: ", accountBalance)
		} else if option == 5 {
			fmt.Println("Thank you for using Go Bank!")
			break
		} else {
			fmt.Println("Invalid Option!")
		}
	
	}

	fmt.Println("Thank you for using Go Bank!")
}