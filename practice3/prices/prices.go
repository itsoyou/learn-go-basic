package prices

import (
	"fmt"

	"example.com/tax-calculator/conversion"
	"example.com/tax-calculator/iomanager"
)

// define a struct
type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"` // field is excluded for json file
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	// map is a key-value pairs!
}

// why pointer? cuz we want to edit the actual value of variable, not the copied one.
func (job *TaxIncludedPricesJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

// function with receiver argument.
func (job *TaxIncludedPricesJob) Process() error {
	err:= job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string) // make() creates empty map

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job) // this already returns error
}

// constructor
func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
