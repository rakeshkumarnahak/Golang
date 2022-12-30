package main

import "fmt"

/***************** IN GO LANG WE CAN DECLARE CONSTANTS *********************/
const LoginToken = "gjisunefv" //The variable whose name starts with a capital letter is a public variable and can be used anywhere throughout the project

func main() {
	var username string = "Rakesh"
	fmt.Println(username)
	fmt.Printf("Type of the variable is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of the variable is: %T \n", isLoggedIn)

	var smallNum uint8 = 255 //8 bit unsigned integer variable
	fmt.Println(smallNum)
	fmt.Printf("Type of the variable is: %T \n", smallNum)

	var largeNum uint64 = 3262352352342342342 //64 bit unsigned integer
	fmt.Println(largeNum)
	fmt.Printf("Type of the variable is: %T \n", largeNum)

	var password int = 56454323 //Integers can also be initialised without mentioning the detailed bitsize
	fmt.Println(password)
	fmt.Printf("Type of the variable is: %T \n", password)

	var smallFloat float32 = 255.2322451242 //32 bit floating point integer
	fmt.Println(smallFloat)
	fmt.Printf("Type of the variable is: %T \n", smallFloat)

	var myName = "Rakesh Kuamr Nahak" //If we don't mention the datatype then the lexer automatically determines the datatype by looking at the value and assignes it to the variable
	fmt.Println(myName)
	fmt.Printf("Type of the variable is: %T \n", myName)

	//Default values in the goLang
	var anotherVariable int //the initial value for an int variable is 0, for string the initial value is a space, for boolean variable it is false.
	fmt.Println(anotherVariable)
	fmt.Printf("Type of the variable is: %T \n", anotherVariable)

	//We can create a variable without writing the var keyword
	//BUT THIS IS ONLY APPLICABLE INSIDE A METHOD ONLY
	numberOfUsers := 80000 // := is the shorthand for declaring anf iniitalising the variable at the same time
	fmt.Println(numberOfUsers)
	fmt.Printf("Type of the variable is: %T \n", numberOfUsers)

	//showing that LoginToken which is publica variable can be used here inside a function also
	fmt.Println(LoginToken)
	fmt.Printf("Type of the variable is: %T \n", LoginToken)

	//Unlike variables, a numeric constant have no datatype untill it is given one
	const number = 11
	fmt.Println(number)
	fmt.Printf("Type of the variable is: %T \n", number)

}
