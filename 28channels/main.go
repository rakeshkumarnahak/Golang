package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	//Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	//Write only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
