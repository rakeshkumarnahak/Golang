package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	//Maps have a key and value related to that key, we can create a map using the make method
	languages := make(map[string]string) //map[key]value

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Map of programming languages:", languages)

	//Declaring and initialising a map in one line
	names := map[string]string{"RKN": "Rakesh Kumar Nahak", "RPN": "Rudra Pratap Nahak", "SM": "Subrajeet Maharana"}
	fmt.Println("Names map", names)

	//Deleting a key value pair in map
	delete(languages, "RB")
	fmt.Println("Languages map after deleting an element:", languages)

	//We can use keys and values individually by getting them inside a loop
	//Here we are using the comma, ok syntax
	for key, value := range languages {
		fmt.Printf("For the key of %v, the value is %v\n", key, value)
	}
}
