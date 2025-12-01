package main

import (
	"fmt"
)

func main() {
	// var investmentAmount float64
	// var expectedReturnRate float64
	// var years float64

	// fmt.Scan(&investmentAmount)
	// fmt.Scan(&expectedReturnRate)
	// fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// fmt.Printf("futureValue: %v\n", futureValue)

	name := "Onat"
	age := 23

	fullStr := fmt.Sprintf("My name is %s and I am %d years old", name, age) // formatting string

	fmt.Println(fullStr)

	fmt.Printf("sum(12, 21): %v\n", sum(12, 21))
}

func sum(a, b int) int {
	return a + b
}
