package main

import "fmt"

// variable in global scope
var global = 200

const CONST_DTA = "SOME CONST DATA"

func main() {

	//declaring variable
	var a = "initial" // it will be infrared to string type
	fmt.Println(a)

	//declaring multiple variables
	var b, c int = 1, 2 // two variables in one line ( and in the second variable we are specifying the type
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// creating variable with no value
	//var k int //default value for int is gonna be zero
	//fmt.Printf(k)

	//shorthand to declare and initializing a variable
	g := "apple"
	fmt.Println(g)

	//constant
	const BASE_URL = "https"

}
