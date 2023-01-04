package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This need to go in a file"

	//Creating the file
	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	//Eriting into the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length of the file is:", length)
	defer file.Close() //file shoul be closed at the end of the operation

	readFile("./mygofile.txt")

}

// Function to read the file
func readFile(filename string) {
	//when any file in the internet is read, it is not read in the string format but it is read in the databyte format
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("The data inside the file is: \n", databyte)	//This will show the result in databytes
	fmt.Println("The data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
