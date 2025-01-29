package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 // value is not going to change in the future

func main() {

	// Different ways to assign variable

	// var investmentAmount float64 = 1000
	// investmentAmount := 1000 (accept suggeseted type)
	// var investmentAmount float64 (default value of float64 is 0.0)
	// investmentAmount (This will not work. MUST need type)

	// Multiple variables assigned in one line

	// var investmentAmount, years float64 = 1000, 10
	// var investmentAmount, years = 1000, "10"
	// var investmentAmount string, years float64 = 1000, 10 (multiple types are not supported)
	// investmentAmount, years := 1000, 10


	var investmentAmount float64 // default value of float64 is 0.0
	var years float64
	expectedReturnRate := 5.5

	// Get User Input using pointer
	//fmt.Print("Investment Amonut: ")
	outputText("Investment Amonut: ")
	fmt.Scan(&investmentAmount) // This is only valid for single-word input

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)



	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// inflation adjusted value
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculatedFutureValues(investmentAmount, expectedReturnRate, years)

	// Sprintf() not only print out but also returns a value
	formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future Value with inflation: %.1f", futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %v\n", futureValue) // v for general value
	// fmt.Printf("Future Value: %f\n", futureValue)
	// fmt.Printf("Future Value: %.0f\n", futureValue) // 0 precision value
	// fmt.Printf("Future Value: %.1f\n", futureValue) // 1 precision value
	// fmt.Println("Future Value with inflation: ", futureRealValue)
	fmt.Print(formattedFutureValue, formattedFutureRealValue)

	// `` backtick allows the linebreak in String. Here the text appears as it is. No more \n
// 	fmt.Printf(`Future Value: %.1f
// Future Value with inflation: %.1f`, futureValue, futureRealValue)
}

// Input Params can also share type annotation
// func outputText(text, text2 string) {
// 	fmt.Print()
// }

func outputText(text string) {
	fmt.Print(text)
}

// func calculatedFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	frv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, frv
// }

func calculatedFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}