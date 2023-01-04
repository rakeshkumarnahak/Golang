package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to math in golang")

	// var myNumOne int = 2
	// var myNumTwo float64 = 4.5

	// fmt.Println("The sum is:", myNumOne+int(myNumTwo))
	// //Here we got the value in integer but lost the precision of 0.5

	// //Generating Random number
	// fmt.Println(rand.Intn(5)) //It'll generate random numbers between 0 and 5(excluding 5)
	// //In most of the modern programming languages such as Go, Python the range is not inclusive.

	// //But it is seen that this doesn't really generate very random numbers because of the seed. We need to provide the seed some random number so that the rand.Intn will generate numbers more randomly.
	// rand.Seed(50)
	// fmt.Println(rand.Intn(5)) //This is also not very random as the randomness revolves around the seed

	// //So wee need to provide seed a number that itself changes everymoment i.e time(nanosecond value)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//Generating random number using crypto
	myNewRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myNewRandomNum)

}
