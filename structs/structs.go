package main

import (
	"fmt"
	"structs-example/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// this is called struct literal notation
	// appUser := user{
	// 	firstName:  userFirstName,
	// 	lastName:   userLastName,
	// 	birthdate:  userBirthdate,
	// 	created_at: time.Now(),
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "1234")

	admin.OutputUserDetailsMethod()
	admin.ClearUserName()
	admin.OutputUserDetailsMethod()

	// alternative ways to define struct literal notation

	// this will be null
	// appUser = user{}

	// if order is same as defined struct
	// if not pass any value then it will be null value.

	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	outputUserDetails(*appUser)
	outputUserDetailswithPoiner(appUser)

	// calling struct's method
	appUser.OutputUserDetailsMethod()
	appUser.ClearUserName()
	appUser.OutputUserDetailsMethod()
}

func outputUserDetails(usr user.User) {
	fmt.Println(usr.FirstName, usr.LastName, usr.Birthdate, usr.Created_at)
}

func outputUserDetailswithPoiner(usr *user.User) {
	// We can access fields of a struct pointer in two ways:
	// 1. Explicitly dereferencing the pointer: (*usr).firstName
	// 2. Using Go's automatic dereferencing: usr.lastName

	fmt.Println((*usr).FirstName, usr.LastName, usr.Birthdate, usr.Created_at)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
