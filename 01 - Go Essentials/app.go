package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10 // only one type declaration per line
	expectedReturnRate := 5.5
	const inflationRate float64 = 2.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
