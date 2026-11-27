package main

import "fmt"

func Banking() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
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
			return
		}

		accountBalance += depositAmount
		fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter your withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Println("You can't withdraw more than you have")
			return
		}

		accountBalance -= withdrawalAmount
		fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
