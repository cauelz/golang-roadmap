package main

import (
	"bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {

	var accountBalance, err = fileops.GetFloatFromFile("balance.txt")

	if err != nil {
		fmt.Println("Error getting balance from file!")
		fmt.Println(err)
		fmt.Println("---------")
	}	
	fmt.Println("Dummy Data: ", randomdata.FullName(randomdata.RandomGender))
	fmt.Println("Welcome to the Go Bank!")
	for {

		presentOptions()
	
		var option int
		fmt.Scanln(&option)
	
		switch option {

		case 1:
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
			fileops.WriteFloatToFile(accountBalance, "balance.txt")
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
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Withdrawal Successful!")
			fileops.WriteFloatToFile(accountBalance, "balance.txt")
			
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