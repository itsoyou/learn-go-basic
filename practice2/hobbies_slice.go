package main

import "fmt"

type Product struct {
	id int
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"CrossFit", "Reading", "Running"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	sliceHobbies := hobbies[:2]
	fmt.Println(sliceHobbies)

	// 4)
	fmt.Println(cap(sliceHobbies))
	sliceHobbies = sliceHobbies[1:3]
	fmt.Println(sliceHobbies)

	// 5)
	courseGoals := []string{"Golang", "Java Spring"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "New Goal"
	courseGoals=append(courseGoals, "AWS Certificate")
	fmt.Println(courseGoals)

	// 7)
	products := []Product{
		{1, "happy", 39.99},
		{2, "day", 19.99},
	}
	products = append(products, Product{3, "holy", 18.45})
	fmt.Println(products)
}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.