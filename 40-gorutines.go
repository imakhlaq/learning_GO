package main

import (
	"fmt"
	"time"
)

func main() {

	var counter int

	for i := 0; i < 1000; i++ {
		//creating go routine
		go func() {
			counter++
		}()
	}
	//sleeping
	time.Sleep(4*time.Second)

	fmt.Println("Expected", 10)
	fmt.Println("Got", counter)
}