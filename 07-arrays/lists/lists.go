package lists

import "fmt"

func main() {
	var productNames [4]string

	fmt.Println(productNames)

	productNames = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	fmt.Println(prices[2])

	featuredPrices := prices[1:4] //including first value, but second value is excluded.
	fmt.Println(featuredPrices)
	// NEGATIVE index does not works. Highest value is the length of list.

	// modifying element in slices means you are modifying it in original array.
	featuredPrices[0] = 14.24
	fmt.Println(prices)

	// slices are not copied from original memory.
	// it's one array memory space, slice is just referencing part of it. efficient!


	fmt.Println(len(prices))
	fmt.Println(cap(prices)) // capacity and length is same here.

	highlightedPrices := featuredPrices[:1]
	fmt.Println(len(highlightedPrices))
	fmt.Println(cap(highlightedPrices)) // Original array of this slice has 3 items
	// capacity can't go further to the starting point(left),
	// but it can go further to the end(right)

	highlightedPrices = highlightedPrices[:3] // we can re-slice highlightedPrices!! O_O
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))


	fmt.Println("==================")

	// Dynamic Array. (No size defined)
	numbers := []float64{10.99, 8.99}
	fmt.Println(numbers[0:1])
	updatedNumbers := append(numbers, 5.99) // does not add to original array, rather creates new array with appended items.
	fmt.Println(updatedNumbers, numbers)

	numbers = append(numbers, 7.99, 5.76, 103.4) // To append to original array, simply re-assign the array. but same thing happens in memory-wise.
	fmt.Println(updatedNumbers, numbers)

	numbers = numbers[1:] // removing first element

	// Remove the ith element
    // numbers = append(numbers[:i], numbers[i+1:]...)

	discountNumbers := []float64{103.2, 10.24, 94.25}

	// numbers = append(numbers, discountNumbers) this throws error, because []float64 isn't float64
	numbers = append(numbers, discountNumbers...)
	// these three dots "..." means take out all the elements and add them separate.

}