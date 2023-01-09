package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()

	r := mux.NewRouter()                       //Creating a new router
	r.HandleFunc("/", SaveHome).Methods("GET") //Making a router run a specific function on a certain path

	log.Fatal(http.ListenAndServe(":4000", r)) //Creating a server in golang and wrapping it around log.Fatal() which would show the error of there is any
}

func greeter() {
	fmt.Println("Hello there mod users")
}

func SaveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}
