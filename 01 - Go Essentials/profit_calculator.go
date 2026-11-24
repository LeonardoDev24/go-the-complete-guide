package main

import "fmt"

func Calculate() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	// Earnings before tax
	var ebt float64 = revenue - expenses
	var profit float64 = ebt * (1 - taxRate/100)
	var ratio float64 = ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Print(ratio)
}
