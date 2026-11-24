package main

import "fmt"

func Calculate() {
	var revenue float64 = GetUserInput("Revenue: ")
	var expenses float64 = GetUserInput("Expenses: ")
	var taxRate float64 = GetUserInput("Tax rate: ")

	calculateFinantials(revenue, expenses, taxRate)
}

func GetUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinantials(revenue, expenses, taxRate float64) {
	var ebt float64 = revenue - expenses
	var profit float64 = ebt * (1 - taxRate/100)
	var ratio float64 = ebt / profit

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
