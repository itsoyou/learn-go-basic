package user

import (
	"errors"
	"fmt"
	"time"
)

// Struct is a custom user type, that can not only store data but also functions.
// Uppercase allows to use the struct in other packages.
type User struct {
	FirstName string
	lastName string // This lowercase field is not accessible from outside of the package.
	birthDate string
	createdAt time.Time
	// Make it lowercase here and use Constructor method!
} // if one of the value is not given, it will be just null


// Struct Embedding: building a struct with existing struct
type Admin struct {
	email string
	password string
	User // Embedding User struct
}

type Sample struct {
	email string
	password string
	Special User // Embedding User struct, if lowercase, hidden field.
}


// (u user) is called "Receiver". Here *user is not mandatory, but just for consistency.
func (u *User) OutputUserDetail() {
	fmt.Println(u.FirstName, u.lastName, u.birthDate)
}


// To edit the value of the variable, you have to hand over the pointer!!
// The function with (u user) will not change the value of the u.
// Because what was passed to the function is a copy of the value.
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

// just a pattern. constructor function is normally named as "new".
// To avoid the creation of copy, we return pointer value!!
func New(firstName, lastName, birthDate string) (*User, error) {

	// validation steps can be added
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}
	// returning pointer so don't create a copy
	return &User{
		FirstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			FirstName: "Admin",
			lastName: "Last",
			birthDate: "11/10/1975",
			createdAt: time.Now(),
		},
	}
}

func NewSample(email, password string) Sample {
	return Sample{
		email: email,
		password: password,
		Special: User{
			FirstName: "Admin",
			lastName: "Last",
			birthDate: "11/10/1975",
			createdAt: time.Now(),
		},
	}
}