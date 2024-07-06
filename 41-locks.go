package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var counter int
	var lock sync.Mutex

	for i := 0; i < 1000; i++ {
		//creating go routine
		go func() {
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}
	//sleeping
	time.Sleep(4 * time.Second)

	fmt.Println("Expected", 10)
	fmt.Println("Got", counter)
}
