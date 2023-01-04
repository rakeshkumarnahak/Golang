package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	//Declaring a slice to print using loop
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//If we want to loop through the whole variable then we can use this syntax

	// for i := range days {
	// 	fmt.Println(days[i]) //But in this syntax we'll get the index in i not the values of days
	// }

	//In the range syntax of loop we'll get two values first is the index and the second is the values of the variable we are looping through

	// for index, value := range days {
	// 	fmt.Printf("Day-%v of the week is %v\n", index, value)
	// }

	//If we don't need any of the value then we can use _
	for _, value := range days {
		fmt.Println("Day is:", value)
	}

	randomValue := 1

	for randomValue < 10 {
		if randomValue == 6 {
			goto banner
		}
		if randomValue == 2 {
			randomValue++
			continue //continue skips one value and continues to next
		} else if randomValue == 7 {
			break
		}
		fmt.Println("Value is:", randomValue)
		randomValue++
	}

	//Declaring a label
banner:
	fmt.Println("Welcome to The Analysers!")

}
