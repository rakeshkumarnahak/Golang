package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//There is mo concept of class in golang, but structs can be used as a replacement
	//There is no inheritance in golang, There is no super or parent in golang

	//Initialising the struct
	rakesh := User{"Rakesh", "Rakesh@go.dev", true, 22}
	fmt.Println(rakesh)
	fmt.Printf("Rakesh details are: %+v\n", rakesh) //To get details with property name we should use %+v or else use %v
	fmt.Printf("User name is %v, and email is %v\n", rakesh.Name, rakesh.Email)

}

//Here the struct name and the properties are exported and they can be used anywhere in the program so they are public packages. That's why we should be name the first letter capital.
type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}
