package main

import "fmt"

//a function that takes two int and return an int

// go requires the explicit return
func sum(num1 int, num2 int) int {

	return num1 + num2
}

// function having similar types of parameter
func addSom(a, b, c int) int {
	return a + b + c
}

// returning multiple values from a function
// (int, string) represent the multiple values u want to return from a function.
// this pattern is used to return error and value from a function
func multiReturn() (int, string) {
	return 1, "ball"
}

//variadic functions can be called with any amount of params. (args in java, rest in js)

func sumK(par ...int) int {

	total := 0

	for index, element := range par {

		fmt.Println("working on ", index)
		total += element
	}
	return total
}

func main() {

	ans := sum(3, 8)

	fmt.Println(ans)

	//taking second value and leaving first returned value
	_, str := multiReturn()

	fmt.Println(str)

	// to call variadic function u can pass any amount of values
	sumK(1, 2, 3, 3, 4, 4, 555)

	// u can also pass an array
	arr := []int{1, 2, 3, 5, 4, 5}

	//spreading array in individual elements
	sumK(arr...)

}
