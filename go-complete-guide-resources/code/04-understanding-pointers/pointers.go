package main

import "fmt"

func main() {
	age := 40

	agePtr := &age

	fmt.Println(age)     // regular value
	fmt.Println(&age)    // pointer of the value
	fmt.Println(agePtr)  // pointer of the value
	fmt.Println(*agePtr) // value of the pointer

	adultYear := getAdultYear(&age)
	fmt.Println(adultYear)
	fmt.Println(age)
}

func getAdultYear(age *int) int {
	*age = *age - 18
	return *age
}
