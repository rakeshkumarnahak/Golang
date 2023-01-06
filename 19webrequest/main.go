package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "http://lco.dev"

func main() {
	fmt.Println("Welcome to web request in golang")

	//Making a http web request
	response, err := http.Get(url)
	checkingNilErr(err)

	fmt.Printf("Response is of type %T\n", response)

	//It is the caller's responsibility to close the connection, and we have to do that at the end of the function so we are using defer
	defer response.Body.Close()

	//Reading the response got from the http call
	databytes, err := ioutil.ReadAll(response.Body)
	checkingNilErr(err)

	content := string(databytes)
	fmt.Println(content)

}

func checkingNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
