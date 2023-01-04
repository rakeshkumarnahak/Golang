package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 8)
	fmt.Println("Result is:", result)

	proResult := proAdder(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Pro result is:", proResult)

	//While calling a multiple return value function we need to mention the exact number of variables as of the no of returned values
	totalMarks, messege := markChecker(56, 34, 23, 89, 54)
	fmt.Printf("Total score is %v and %v", totalMarks, messege)

}

func greeter() {
	fmt.Println("Welcome to functions in golang")
}

//We have to mention what kind of value the function is returning i.e the datatype of the returned value after the arguments, this is called "Function Signature"
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//When we don't the exact number of arguments we can take any no of values we want using the below syntax, these type of functions are caleld "Variadic Functions"
//In variadic functions all the arguments are stored inside a slice which can be used further
func proAdder(values ...int) int {
	var result int
	for index := range values {
		result += values[index]
	}
	return result
}

//In functions of golang we can return multiple values, these are called multiple return value functions, but in that case we'll need same number of variables as of the returned values to store them  while calling the function
func markChecker(marks ...int) (int, string) {
	var totalMarks int
	var messege string
	for _, mark := range marks {
		totalMarks += mark
	}

	if totalMarks >= 300 {
		messege = "You scored fantastic"
	} else if totalMarks >= 200 && totalMarks < 300 {
		messege = "You scored well"
	} else if totalMarks >= 100 && totalMarks < 200 {
		messege = "Your score was average"
	} else {
		messege = "You need more practice"
	}

	return totalMarks, messege
}
