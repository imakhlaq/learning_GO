package main

import (
	"fmt"
	"time"
)

func worker1(c chan<- string) { //only writable channel
	time.Sleep(2 * time.Second)
	//writing to the channel
	c <- "hello mr channel"
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go worker1(c1)
	go worker1(c2) //they are running concurrently so the total wait time is 2 sec
	go worker1(c3)

	for i := 0; i < 3; i++ { //three times because we are waiting for three channels
		//select is a blocking operation
		select { //use select to await both of these values simultaneously, printing each one as it arrives.
		case data := <-c1:
			fmt.Println(data)
		case data := <-c2:
			fmt.Println(data)
		case data := <-c3:
			fmt.Println(data)
		}
	}
}
