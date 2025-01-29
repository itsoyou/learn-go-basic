package fileops
// To be a package, a file must be under a separate directory than a main package.
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// function name should be capitalized to be exported from the external package
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value) // this gives string value of number
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // Readfile returns 2 values, data and error
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err :=strconv.ParseFloat(valueText, 64) // PraseFloat also returns 2 values.
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}