package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 // explicit type declaration
	expectedReturnRate := 5.5           // declare and assign a variable where the type should be inferred by Go
	var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
