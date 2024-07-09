package main

import (
	"fmt"
	"sync"
)

type person1 struct {
	name string
	age  int
}

func main1() {

	var waitGroup sync.WaitGroup

	p := person1{name: "Akhlaq", age: 23}

	for range 1000 {
		waitGroup.Add(1)
		go func() {
			p.age = p.age + 1 //multiple goroutines are trying to update age
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println(p.age) // multiple goroutine try to access and update age it can create race condition
}

func main() {

	//its better to have mutex inside the struct
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex //mutex to lock down the shared resource

	p := person1{name: "Akhlaq", age: 23}

	for range 1000 {
		waitGroup.Add(1)
		go func() {
			mutex.Lock()      //locking the shared resource so no other goroutine can access in same time
			p.age = p.age + 1 //multiple goroutines are trying to update age
			mutex.Unlock()    //unlocking the shared resource so other goroutine can access in same time
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println(p.age)
}
