package main

import "fmt"

func main() {

	c := make(chan string, 100)

	c <- "data"
	c <- "ups"
	c <- "net"

	close(c)
	//close is necessary because channel have 100 buffer and range will try to read all 100 buffer.
	//and hence no one is writing on the channel and it will block the goroutine.
	// And golang will detect it and panic with deadlock

	// for using range u have to close the channel.
	//range will try to keep on reading from channel and non-one writing to the channel so goroutine will be blocked
	// go will detect it and caused panic with deadlock
	for data := range c {
		fmt.Println(data)
	}

	// message,ok:= <- c
	//check if ok is nil. then channel have no data
}
