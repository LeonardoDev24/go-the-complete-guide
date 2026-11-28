package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func Banking() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit\n")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
		} else if choice == 3 {
			fmt.Print("Enter your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have")
				// return
				continue
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing our bank!")
}

func writeBalanceToFile(balance float64) {
	var balanceText string = fmt.Sprint(balance)
	var balanceByte []byte = []byte(balanceText)
	os.WriteFile(accountBalanceFile, balanceByte, 0644)
}

func getBalanceFromFile() float64 {
	var data, _ = os.ReadFile(accountBalanceFile)
	var balanceText string = string(data)
	var balance, _ = strconv.ParseFloat(balanceText, 64)
	return balance
}
