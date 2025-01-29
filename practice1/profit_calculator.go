package main

// Goals
// 1) Validate user user input
//   -> show error message & exit if invalid input is provided
//		negative or 0
// 2) Store calculated results into file

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const profitCalculatorFile = "profit.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		log.Fatal(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		log.Fatal(err)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		log.Fatal(err)
	}

	ebt, profit, ratio := calculateFigures(revenue, expenses, taxRate)
	writeResultToFile(ebt, profit, ratio)

	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}

// Using pointer as parameter
// func getUserInput(infoText string, destination *float64) {
// 	fmt.Print(infoText)
// 	fmt.Scan(destination)
// }
// getUserInput("Revenue: ", &revenue)
// getUserInput("Expenses: ", &expenses)
// getUserInput("Tax Rate: ", &taxRate)


func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <=0 {
		return 0, errors.New("number can't be negative or 0")
	}
	return userInput, nil
}

func calculateFigures(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (revenue * taxRate/100)
	ratio = ebt/profit
	return ebt, profit, ratio
}

func writeResultToFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(profitCalculatorFile, []byte(result), 0644)
}