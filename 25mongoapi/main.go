package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rakeshkumarnahak/mongoapi/router"
)

func main() {
	r := router.Router()

	fmt.Println("Server is Starting")

	log.Fatal(http.ListenAndServe(":5000", r))

	fmt.Println("Server is running at port 4000")
}
