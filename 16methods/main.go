package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")
	//There is mo concept of class in golang, but structs can be used as a replacement
	//There is no inheritance in golang, There is no super or parent in golang

	//Initialising the struct
	rakesh := User{"Rakesh", "Rakesh@go.dev", true, 22}
	fmt.Println(rakesh)
	fmt.Printf("Rakesh details are: %+v\n", rakesh) //To get details with property name we should use %+v or else use %v
	fmt.Printf("User name is %v, and email is %v\n", rakesh.Name, rakesh.Email)
	rakesh.getVerificationStatus()
	rakesh.NewEmail()
	fmt.Printf("User name is %v, and email is %v\n", rakesh.Name, rakesh.Email) //For the response to this command you can see that we get email as befor because inside NewEmail() method a copy of the Email value is passed so the original value is remained unchanged.
	//To change the original value of the Email we have to pass the reciever type in pointer
	rakesh.NewAge()
	fmt.Printf("User name is %v and age is %v\n", rakesh.Name, rakesh.Age) //Here you can see the age is also changed outside of the function call as well. This happed because of the pointer reciever type
	fmt.Printf("Rakesh details are: %+v\n", rakesh)
}

//Here the struct name and the properties are exported and they can be used anywhere in the program so they are public packages. That's why we should be name the first letter capital.
type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

//Functions that are passed into a class or struct in golang is called a method
// Every method has a reciever type in which any struct is passed whose part is that method is, this can be done by passing a value or by passing a pointer of that struct, below is the example of the value reciever type
//func (receiverType) functionName() {}
func (u User) getVerificationStatus() {
	fmt.Println("Is user verified:", u.Verified)
}

//If you want to export the method then write the first letter in capital
func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("The new email is:", u.Email)
}

//Method having the pointer reciever type
func (u *User) NewAge() {
	u.Age = 20
	fmt.Println("The new age is:", u.Age)
}
