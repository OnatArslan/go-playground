package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Println("Waiting for your choice")
		fmt.Scan(&choice)

		fmt.Println("Your choice is:", choice)

		if choice == 1 {
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		} else if choice == 2 {

			var amount float64
			fmt.Println("Please enter an amount for deposit")
			fmt.Scan(&amount)

			if amount >= 0 {
				accountBalance += amount
				fmt.Printf("Your new balance is %.2f\n", accountBalance)
			} else {
				fmt.Println("Negative amount given")
				break
			}
		} else if choice == 3 {

			var amount float64
			fmt.Println("Please enter an amount for withdraw")
			fmt.Scan(&amount)
			if accountBalance >= amount {
				accountBalance -= amount

				fmt.Printf("Your new account balance is %.2f", accountBalance)
			} else {
				fmt.Println("You don't have that kind of money")
				break
			}
		} else {
			break
		}
	}

}
