package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang")

	//Declaring an array
	var fruitList [4]string //It is compulsury to mention the size of the array and the datatype of its elements

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Length of the fruit lsit is:", len(fruitList)) //Here the length will come as 4 not 3(no o elements inside the array)
	//Length of the array is same as it is mentioned during the declaration
	//In array we don't get to know the no of elements in the array but we get the size of the array during its declaration

	//Another way to declare arrays
	var vegList = [5]string{"Potato", "Beans", "mushroom"}
	fmt.Println("Veg list is:", vegList)
	fmt.Println("length of the veg list is:", len(vegList))

	//NOTE: In many programming languages arrays are more widely used but in Golang arrays don't have many usecases. Slices are more frequently used here.
}
