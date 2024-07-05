package main

import (
	"fmt"
)

func produce(message chan string) {
	message <- "ping" //writing to the channel so other go routines can consume it
}

func main() {

	//creating a channel. channels are typed by the values they convey
	message := make(chan string)

	//passing the channel to a go routine
	go produce(message)

	//anonyms goroutine
	go func() {
		//reading from channel
		msg := <-message
		fmt.Println(msg)

		//writing to the channel
		message <- "supper"
	}()
	// reading from channel that is written by second goroutine
	msg := <-message

	fmt.Println(msg, "from main")

	close(message) //closing the channel
}

// This will cause a deadlock. Because first go routine write to the channel and second goroutine read from it and does't write to it.
// And we again try to read from the channel that cause the deadlock.
func mainDeadlock() {

	message := make(chan string)
	go produce(message)

	go func() {
		msg := <-message
		fmt.Println(msg)
	}()

	msg := <-message

	fmt.Println(msg, "from main")
}
