package main

import (
	// "example.com/tax-calculator/cmdmanager"
	"fmt"

	"example.com/tax-calculator/filemanager"
	"example.com/tax-calculator/prices"
)

func main() {

	// taxRates := []float64{0, 0.07, 0.1, 0.15}
	// for _, taxRate := range taxRates {
	// 	iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
	// 	// iom := cmdmanager.New()

	// 	priceJob := prices.NewTaxIncludedPricesJob(iom, taxRate)
	// 	err := priceJob.Process()
	// 	if err != nil {
	// 		fmt.Println("error!")
	// 		fmt.Println(err)
	// 	}
	// }

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {

		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPricesJob(iom, taxRate)

		// goroutines don't support return values.
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select { // Select is special statement for channels.
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done!")
		}
	}

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
