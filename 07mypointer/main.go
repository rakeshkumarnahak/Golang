package main

import "fmt"

func main() {
	//Initialising the pointer variable
	// var ptr *int
	// fmt.Println("value of the pointer variable is:", ptr) //The value inside the pointer is nil

	myNumber := 23
	var ptr = &myNumber //Here the reference of myNumber is sotred in ptr (referece means '&')

	fmt.Println("value of pointer is:", ptr)             //ptr will show the reference of the variable ptr
	fmt.Println("value of the actual pointer is:", *ptr) // "*ptr" shows the value stored inside the variable to which ptr is pointing to

	*ptr = *ptr + 2
	fmt.Println("Value of the actual pointer after addition:", *ptr)
	fmt.Println("New value of the main myNumber variable is:", myNumber)
}
