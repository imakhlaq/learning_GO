package main

import "fmt"

/*

Go have pointers, it allows to pass references of values. i.e memory address.
Go does not have pointer arithmetic.

to create pointer variable you use  " var a *int = & value "; => "a" is a going to store a pointer

to reference the address of a variable you use ===> &num

if u have a pointer variable and u want to get its value use *num

*/

func subValue(num int) int {
	num -= 1
	return num
}

func subValueptr(num *int) int {

	//||==> to tell its a pointer variable
	//      ||==> referencing pointer variable value
	*num = *num - 1

	return *num
}

func main() {

	var num = 10

	fmt.Println("Initial ", num)
	subValue(num)
	fmt.Println("subValue", num)
	subValueptr(&num)
	fmt.Println("subValueptr", num)

	//===========================================================================================================

	var k *int = &num //k have pointer to num variable

	fmt.Println(*k, "to get the value of pointer")
	fmt.Println(&num, "to get the address of a variable")
}
