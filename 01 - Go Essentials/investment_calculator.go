package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func CalculateInvestment() {
	var investmentAmount float64 = GetUserInput("Investment amount: $")
	var expectedReturnRate float64 = GetUserInput("Expected return rate: ")
	var years float64 = GetUserInput("Years: ")

	var futureValue, futureRealValue float64 = calculateFutureValues(investmentAmount, expectedReturnRate, years)

	var formattedFutureValue string = fmt.Sprintf("Future value: %.2f\n", futureValue)
	var formattedFutureRealValue string = fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFutureValue, formattedFutureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	var fv float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var frv float64 = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
