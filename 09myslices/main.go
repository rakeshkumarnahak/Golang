package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	//Declaring and initialising slices
	//There are two ways we can declare slices

	var fruitList = []string{"Apple", "Banana", "Peach"}
	//NOTE: In arrays we had to mention the length of the array while declaring it but in case of slice we don't need to mention the size, which means in slices we can add as many elements we want.
	//In this way of declaring the slice we have to initialise it.

	fmt.Println("Fruit list slice:", fruitList)

	//In slices we can add elements later using append function
	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println("Fruit list slice after adding more fruits", fruitList)

	//Slices has a special slice operator

	//This slices the fruitList slice into new one having elements from index 1 to index 3(excluding 3)
	fruitList = fruitList[1:3] //syntax: sliceName(startingIndex : endingIndex)  (Note that this will not include the element at the finalIndex )
	fmt.Println("Fruit list slice after slicing:", fruitList)

	// fruitList = fruitList[1:] //This will make a new slice having elements from index 1 to end
	// fmt.Println("Fruit list slice after slicing:", fruitList)

	// fruitList = fruitList[:4] //This will make a new slice having elements starting from begining to the index 4(excluding 4th index element)
	// fmt.Println("Fruit list slice after slicing:", fruitList)

	//Another way of declaring Slices
	//We can use the make() function to create a memory for a slice
	highScores := make([]int, 4) //Here 4 show the size of the slice(we need to mention the size in make method)
	highScores[0] = 453
	highScores[1] = 786
	highScores[2] = 346
	highScores[3] = 768
	// highScores[4] = 444	//Here this will show error because it is out of range
	//NOTE: But in slices we can add as much elements we want right so why it is showing out of range?

	fmt.Println("High scores slice:", highScores)

	// We can elements to sices using the append mehod
	highScores = append(highScores, 675, 855, 897)
	fmt.Println("High scores slice after adding more elements to it:", highScores)

	//We can check whether a slice is sorted or not
	fmt.Println("Is the high scores slice sorted:", sort.IntsAreSorted(highScores))

	//We can sort slices using sort package
	sort.Ints(highScores)
	fmt.Println("Sorted high scores slice", highScores)

	fmt.Println("Is the high scores slice sorted:", sort.IntsAreSorted(highScores))

	/***************** We can copy one slice into another ********************/
	newHighScores := highScores[1:4] //Making an slice having highscores from index 1 to 3

	var finalScore = make([]int, 2) //Making an empty slice
	fmt.Println("Final scores before copying:", finalScore)

	//copying newHighScores slice into finalScores slice
	copy(finalScore, newHighScores) //Copying syntax should be like copy(destination, source)
	fmt.Println("Final scores after copying ", finalScore)

	/*************** How to remove an element from slice based on index ****************/
	progLanguages := []string{"javascript", "ruby", "swift", "C", "C++"}
	fmt.Println("Programming languages slice:", progLanguages)

	var index int = 2 //saving the index (of the element that is to be deleted)in a variable

	//Removing the element at index 2(value inside the variable)

	//Here the append method reallocated the memory with a new slice which is the combination of two slices.
	//First is the slice having elements from begining to the mentioned index(excluding the element at the index)
	//second is the slice having elements from the index+1 to the end
	progLanguages = append(progLanguages[:index], progLanguages[index+1:]...)
	fmt.Println("Programming languages slice after deleting the element of index 2:", progLanguages)
}
