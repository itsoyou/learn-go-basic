package fnarevalues

import "fmt"

type transformFn func(int, []string) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFn(&numbers)
	transformerFn2 := getTransformerFn(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(transformedMoreNumbers)

}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// func transformNumbers2(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

func getTransformerFunction() func(int) int {
	return double
	// return double()
	// This will return the result of the function execution
}

func getTransformerFn(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}