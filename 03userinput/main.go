package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating for our please")

	// comma, ok || comma, err  syntax
	//it should be like [input, error := reader.ReadString('') ] but if we are not going to any of the two variable then we can replace that variable by an underscore.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating ", input)
	fmt.Printf("Type of this rating is %T", input)
}
