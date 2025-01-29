package main

import "fmt"

func main() {
	age := 32 // regular variable

	// normaly name doesn't contains "Pointer"
	var agePointer *int
	agePointer = &age

	// & to get the address. &Variable -> address
	// * to get the value. *Pointer -> value

	// Pointer's null value is called "nil": absence of an address value
	// pointer is poining at no address or no value

	// de-referencing! get actuall value of the pointer's adress
	// agePointer will return 0x1400000e0b0
	// *agePointer will return 32

	fmt.Println("Age:", *agePointer)


	editAgeToAdultYears(agePointer)
	fmt.Println(age)

}

func editAgeToAdultYears(age *int) {
	// with this, age value itself does not change.
	//return *age - 18

	// with this, age value has been changed.
	*age = *age - 18
}

