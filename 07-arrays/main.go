package main

import "fmt"

// type alias.
type floatMap map[string]float64

func (m floatMap)output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // creating empty slice with 2 slots, allocating 5 spaces.
	// if fixed length -> array
	// if non-defined length -> slice

	// make function let Go knows that I am going to add more items into the array
	// Go will create bigger array space, so it doesn't have to re-create it every time

	userNames[0] = "Julie" // userNames := []string{} with this it will throw error.


	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)

	courseRatings :=make(map[string]float64, 3) //make for map only accepts 1 argument.

	courseRatings["go"] = 3.5
	courseRatings["react"] = 4.5
	courseRatings["java"] = 5.0
	fmt.Println(courseRatings)

	// using type alias
	priceTags :=make(floatMap, 2)
	priceTags.output()

	// looping through slices
	for index, value := range userNames {
		fmt.Println(index, ":", value)
	}

	// looping through map
	for key, value := range courseRatings {
		fmt.Println(key, ":", value)
	}
}