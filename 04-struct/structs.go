package main

import (
	"fmt"

	"example.com/struct/user"
)

// "type" keyword can also used for assigning an alias!!
type customInt int

func (number customInt) log(){
	fmt.Println(number)
}

func main() {
	var age customInt = 20
	age.log()


	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// 1. Define a variable using the basic user struct
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// 2. Define a variable using constructor method
	// appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	// 3. Define a varaible via Package using basic User struct
	// appUser = &user.User{
	// 	FirstName: userFirstName,
	// 	lastName: userLastName, // this is not accessible from here because it is lowercase
	// }

	// 4. Define a varaible via Package using Constructor
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Creating empty user
	// appUser2 := user{}

	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetail()
	admin.ClearUserName()
	admin.OutputUserDetail() //admin.User.OutputUserDetail()

	// outputUserDetail(&appUser)
	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()


}

// func outputUserDetail(u user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

// func outputUserDetail(u *user) {
// 	// still works without dereferencing. It is allowed by Go.
// 	// (*u).firstName
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//fmt.Scan(&value) // this allow user to makes multiple empty lines until user input.
	fmt.Scanln(&value) // hitting enter key will just accepted as emtpy string.
	return value
}