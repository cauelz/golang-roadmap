package main

import "fmt"

func run() {

	var accountBalance float64 = 0.0

	for {

		switch option := getOption(); option {
		case 1:
			createAccount(&accountBalance)
		case 2:
			deposit(&accountBalance)
		case 3:
			withdraw(&accountBalance)
		case 4:
			checkBalance(accountBalance)
		case 5:
			return
		default:
			println("Invalid Option!")
		}
	}
}

func getOption() int {
	println("Please select an option:")
	println("1. Create Account")
	println("2. Deposit")
	println("3. Withdraw")
	println("4. Check Balance")
	println("5. Exit")

	var option int
	fmt.Scanln(&option)
	return option
}

func createAccount(accountBalance *float64) {
	fmt.Println("Enter Account Number:")
	var accountNumber int
	fmt.Scanln(&accountNumber)
	fmt.Println("Enter Initial Deposit:")
	var initialDeposit float64
	fmt.Scanln(&initialDeposit)
	*accountBalance = initialDeposit
	fmt.Println("Account Created Successfully!")
}

func deposit(accountBalance *float64) {
	fmt.Println("Enter Deposit Amount:")
	var depositAmount float64
	fmt.Scanln(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Deposit Amount should be greater than 0!")
		return
	}

	*accountBalance += depositAmount
	fmt.Println("Deposit Successful!")
}

func withdraw(accountBalance *float64) {
	fmt.Println("Enter Withdrawal Amount:")
	var withdrawalAmount float64
	fmt.Scanln(&withdrawalAmount)

	if withdrawalAmount <= 0 {
		fmt.Println("Withdrawal Amount should be greater than 0!")
		return
	}

	if withdrawalAmount > *accountBalance {
		fmt.Println("Insufficient Funds!")
		fmt.Println("Your Account Balance is: ", *accountBalance)
		return
	}

	*accountBalance -= withdrawalAmount
	fmt.Println("Withdrawal Successful!")
}

func checkBalance(accountBalance float64) {
	fmt.Println("Account Balance: ", accountBalance)
}