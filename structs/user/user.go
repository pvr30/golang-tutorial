package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName  string
	LastName   string
	Birthdate  string
	Created_at time.Time
}

// struct embedding
type Admin struct {
	email    string
	password string
	User
}

// user struct's method
func (usr User) OutputUserDetailsMethod() {
	fmt.Println(usr.FirstName, usr.LastName, usr.Birthdate, usr.Created_at)
}

func (usr *User) ClearUserName() {
	// We use a pointer receiver here to modify the original struct's data.
	// If we used a value receiver, the method would work on a copy of the struct,
	// and changes would not affect the original instance.
	usr.FirstName = ""
	usr.LastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName:  "Admin",
			LastName:   "Admin",
			Birthdate:  "03/30/2001",
			Created_at: time.Now(),
		},
	}
}

// Constructor Function to create user.
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		FirstName:  firstName,
		LastName:   lastName,
		Birthdate:  birthdate,
		Created_at: time.Now(),
	}, nil
}
