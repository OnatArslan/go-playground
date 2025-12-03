package main

import (
	"fmt"
	"time"

	"example.com/structs/user"
)

func main() {
	// FirstName := getUserData("Please enter your first name")
	// LastName := getUserData("Please enter your last name")
	// BirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	onat := user.User{
		FirstName: "Onat",
		LastName:  "Arslan",
		BirthDate: "29/11/2001",
		CreatedAt: time.Now(),
	}
	// user.FirstName = "Onat"

	fmt.Println(onat)

	mel := user.CreateUser("Melisa", "Ozgur", "22/11/2004")
	fmt.Println(mel.CreatedAt.Format("2006-01-02 15:04:05"))

	onat.SetFirstName("BabaPro")
	fmt.Printf("onat: %v\n", onat)
}
