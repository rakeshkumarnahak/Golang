package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang")

	var result string
	loginCount := 23

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watchout"
	} else {
		result = "Exactly 10 loginCount"
	}

	fmt.Println(result)

	//There is a special syntax of if else in golang
	//Here we declare an initialise a variable and also put a condition around it
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("NUmber is NOT less than 10")
	}
}
