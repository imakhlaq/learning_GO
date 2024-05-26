package main

import (
	"fmt"
	"time"
)

type student struct {
	id   int32
	name string
}

// it can modify the receiving struct (RECOMMENDED)
// use pointer receiver type to avoid method calls or to allow the method to mutate the receiving struct
func (s *student) printName() {

	fmt.Println(s.name, " id is ", s.id)

}

// it cant modify receiving struct
func (s student) calcAge() {
	fmt.Println(s.id, "age is", time.Now().Day())

}

func main() {

	r := student{name: "jami", id: 1111}

	r.calcAge()
	r.printName()

}
