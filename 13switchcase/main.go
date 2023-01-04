package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is ", diceNumber)

	switch diceNumber {
	//Here case 1 doesn't mean case No 1 but the case related the concerning value of diceNumber. Instead of 1 or 2 we can write any other value of diceNumber
	case 1:
		fmt.Println("Dice value is 1 and you can start")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough //Fallthrough is not a default behaviour in switchcase but we can specifcally mention fallthrough where we want it and result of the next will also be shown
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	}
}
