package anonymousfn

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	// use Anonymous function, that doesn't have a name
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

//createTransformer is a "factory function" that creates other function with different configuration.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		// factor value is now locked in this scope. This is closure.
		return number * factor
	}
}

// Every anonymous function is closure