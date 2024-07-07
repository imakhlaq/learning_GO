package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	//external call that gives result after 3 second by writing to the channel
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "result 1"
	}()

	//awaiting for the result
	//only the first channel that sends data will be considered rest will be ignored
	select {
	case res := <-c1:
		fmt.Println(res)
		//if external call does not complete in 2 second we timeout it
	case <-time.After(2 * time.Second):
		fmt.Println("cancelled")
	}
}
