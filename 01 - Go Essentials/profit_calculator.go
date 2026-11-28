package main

import (
	"errors"
	"fmt"
	"os"
)

func CalculateProfit() {
	var revenue, errRev = GetUserInput("Revenue: ")
	var expenses, errExp = GetUserInput("Expenses: ")
	var taxRate, errTax = GetUserInput("Tax rate: ")

	if errRev != nil || errExp != nil || errTax != nil {
		fmt.Println(errExp.Error())
		return
	}

	calculateFinantials(revenue, expenses, taxRate)
}

func GetUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInput, nil
}

func calculateFinantials(revenue, expenses, taxRate float64) {
	var ebt float64 = revenue - expenses
	var profit float64 = ebt * (1 - taxRate/100)
	var ratio float64 = ebt / profit

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	var results string = fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
