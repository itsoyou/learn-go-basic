package main

import (
	// "example.com/tax-calculator/cmdmanager"
	"fmt"

	"example.com/tax-calculator/filemanager"
	"example.com/tax-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// iom := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPricesJob(iom, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("error!")
			fmt.Println(err)
		}
	}

}
