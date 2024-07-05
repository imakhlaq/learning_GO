package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// A goroutine is a lightweight thread of execution.
func main() {

	go f("imakhlaq") //launching go routine

	var counter int

	for i := 0; i < 1000; i++ {
		//launching anonymous go routine
		go func() {
			counter++
		}()
	}

	//sleeping
	time.Sleep(4 * time.Second)

	fmt.Println("Expected", 10)
	fmt.Println("Got", counter)
}
