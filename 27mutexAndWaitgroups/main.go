package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Mutex ad wait groups")

	wg := &sync.WaitGroup{} //Waitgroup variable should be  pointer
	mut := &sync.RWMutex{}  //Pointer

	scores := []int{1}

	//We can add waitgroups one by one or we can add total no of go routines at once
	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Four R")
		mut.Lock()
		scores = append(scores, 4)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.RLock()
		fmt.Println(scores)

		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(scores)
}
