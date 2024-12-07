package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var investmentAmount float64 = 10000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// Different methods to define variables
	// var investmentAmount, years float64 = 10000, 10

	// expectedReturnRate := 5.5

	// investmentAmount, years, expectedReturnRate := 10000.0, 10.0, 5.5

	// Type conversion
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // take the input

	// fmt.Print("Expected Return: ")
	outputText("Expected Return: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Value with adjustment:", futureRealValue)

	fmt.Printf("Future Value: %v\nFuture Value with adjustment: %v\n", futureValue, futureRealValue)

	// multiline strings
	fmt.Printf(`Future Value: %.1f\n
	Future Value with adjustment: %.1f`, futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedFVA := fmt.Sprintf("Future Value with adjustment: %.0f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFVA)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}

// Alternative way to return a value

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
// 	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
// 	return
// }
