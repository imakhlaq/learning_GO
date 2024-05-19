package main

import "fmt"

// go support anonymous function which can from closures.
// anonymous function have no name

func counter() func() int {
	i := 0

	//returning a function
	return func() int {
		//using parent function variable in child because of closures
		i++
		return i
	}
}

func main() {

	nextInt := counter()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
