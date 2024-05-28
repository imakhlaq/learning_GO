package main

import "fmt"

// variable in global scope
var global = 200

const port = "SOME CONST DATA"

// if u have to declare multiple variables
const (
	data   = 10
	father = "some"
	kafka  = "utl"
)

//hahi:="delo" does not work in global scope

func main() {

	//declaring variable
	var a = "initial" // it will be infrared to string type
	fmt.Println(a)

	//declaring multiple variables
	var b, c int = 1, 2 // two variables in one line ( and in the second variable we are specifying the type

	// if u have to declare multiple variables
	var (
		data   = 10
		father = "some"
		kafka  = "utl"
	)

	fmt.Println(b, c, data, father, kafka)

	var d = true
	fmt.Println(d)

	// creating variable with no value
	//var k int //default value for int is gonna be 0
	//fmt.Printf(k)

	//shorthand to declare and initializing a variable
	g := "apple"
	fmt.Println(g)

	//constant
	const base_url = "https"

}
