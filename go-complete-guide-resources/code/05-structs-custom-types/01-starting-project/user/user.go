package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

// Defining one setter method on struct User
func (u *User) SetFirstName(newName string) {
	u.FirstName = newName
}

func GetUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}

func CreateUser(FirstName, LastName, BirthDate string) User {
	user := User{
		FirstName: FirstName,
		LastName:  LastName,
		BirthDate: BirthDate,
		CreatedAt: time.Now(),
	}
	return user
}

// We are not using *u.FirstName because go doing it auto because of readabilty
func LogUserData(u *User) {
	// We can write (*u).FirstName but when we do u.FirstName go do pointer defference
	fmt.Println((*u).FirstName, u.LastName, u.BirthDate, u.CreatedAt)
}
