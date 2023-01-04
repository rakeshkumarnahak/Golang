package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")

	//Defer is very simple. Whenever there is a defer keyword in a line those lines get piled up in the last of the function and they are run after all the code of the function is run.
	//If there are more than one defer statement in a function then they are run in "LIFO" basis i.e "Last In First Out".  All the defer statements run in the reverse order as they were declared

	defer fmt.Println("World!")
	fmt.Println("Hello")
	myDefer() //myDefer is not a defer statement so it'll run after Hello
	//Apparent Result is: Hello 43210 Two One World!

	defer fmt.Println("One")
	defer fmt.Println("Two")
	//Here all the defer statements got piled up in the last of the function and then run in the reverse order of their declaration i.e LIFO(Last In First Out)
	//So the Result is Two One World

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
