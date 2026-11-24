package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	const inflationRate float64 = 2.5

	fmt.Print("Investment amount: $")
	fmt.Scan(&investmentAmount) // pass the pointer to the variable

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	var formattedFutureValue string = fmt.Sprintf("Future value: %.2f\n", futureValue)
	var formattedFutureRealValue string = fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFutureValue, formattedFutureRealValue)
}
