package main

import (
	"fmt"
	"sync"
)

func Worker(num int, wait *sync.WaitGroup) {

	fmt.Println("DOING WORKKK ", num)
	wait.Done()
}

func main() {

	var wait sync.WaitGroup

	for i := range 10 {
		wait.Add(1)
		go Worker(i, &wait) //remember to pass reference/pointer
	}

	wait.Wait() //waiting for all goroutine to finish before finishing main goroutine
}
