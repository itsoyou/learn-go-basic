package main

import "fmt"

func main() {

	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	fmt.Println(sum)

	sum2 := variadicSumup(3, 67, 17, 2, 1)
	fmt.Println(sum2)

	sum3 := variadicSumupWithStartingVal(100, 1, 2, -5)
	fmt.Println(sum3)

	anotherSum := variadicSumup(numbers...) // numbers became 3 separate parameter.
	fmt.Println(anotherSum)
	anotherSum2 := variadicSumupWithStartingVal(200, numbers...)
	fmt.Println(anotherSum2)

}

func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// Variadic functions: work with any amount of parameters.
func variadicSumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func variadicSumupWithStartingVal(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, val := range numbers {
		sum += val
	}
	return sum
}