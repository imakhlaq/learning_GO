package main

import "fmt"

// CustomList generic for creating a list
type CustomList[T any] struct {
	list []T
}

// method to add element to the list
func (c *CustomList[T]) addToList(data T) bool {

	ts := append(c.list, data)
	c.list = ts

	return true
}

// constructor for a newList
func newList[T any]() *CustomList[T] {

	return &CustomList[T]{
		list: make([]T, 0),
	}
}

// generic for func
func someStuff[T any, L any](param1 T, param2 L) T {

	return param1
}

func main() {

	myList := newList[int]()

	myList.addToList(2)
	myList.addToList(4)

	//u cant passe string
	//myList.addToList("dadad")
	fmt.Println(myList.list)

	//generic func
	i := someStuff("data", 1) //types will be inferred
	fmt.Println(i)

}
