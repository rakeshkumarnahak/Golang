package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// creating a waitgroup
var wg sync.WaitGroup //It is a pointer

// Creating a Mutex
// Mutex locks the memory for a single thread to write untill it's done. After the thread completes its job mutex unlocks it and allows the next thread to come in
var mut sync.Mutex

func main() {
	fmt.Println("Welcome to go routines")
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web) //With go keyword it becomes a go routine which takes some time to run so we have to wait. For this we use waitGroups
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	//ending the wait of the waitgroup in the end
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()                          //Locks the memory so only the current go routine can access the memory address
		signals = append(signals, endpoint) //Here all the go routines are accessing this signal memory to append the endpoint at the sametime, which is not good. So here mutex locks one go routine to append first, after it's done it unlocks it and the next go routine is allowed to do its job.
		mut.Unlock()                        //Unlocks the memory for next go routines
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }
