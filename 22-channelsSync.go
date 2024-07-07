package main

import "fmt"

func worker(signal chan bool) {

	for i := 0; i <= 100; i++ {
		fmt.Println("Hello", i)
	}

	//signalling that goroutines is completed
	signal <- true
}

func main() {

	signal := make(chan bool)
	go worker(signal)

	//reading data channel channel.
	//So it will block the main goroutine and wait for the data to be available in channel.
	<-signal
}
