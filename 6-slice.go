package main

import (
	"fmt"
	"slices"
)

func main() {

	//unlike arrays slices only needed to be type by element not by no of elements(means its dynamic array)

	//creating slice is same as an array, but you don't specify the size

	//declaring slice
	var sliceDemo []int
	// and un initializing slice will be nil
	fmt.Println(sliceDemo, sliceDemo == nil)

	//if u want to specify the initial capacity of slice
	sliceDemo = make([]int, 4)

	//adding values to the slice
	for i := 0; i < len(sliceDemo); i++ {
		sliceDemo[i] = i
	}

	fmt.Print(sliceDemo, "length of slice ", len(sliceDemo))

	//if you know the capacity of slice a head of time u can pass it in parameter
	shortSlice := make([]string, 5)

	//shot way of making slice
	shotSlice := make([]bool, 4)
	fmt.Println(shotSlice)

	//slice provide lots of extra features on array like append in a slice
	// append reruns a new slice.
	newSlice := append(shortSlice, "appending", "moreAppending")
	fmt.Println(newSlice)

	//u can copy one slice into another
	c1 := make([]string, 3)

	c1[1] = "need to copy"

	c2 := make([]string, 3)

	//coping one slice into another
	copy(c2, c1)

	fmt.Println(c2)

	//slices support "slice" operations eg str.slice()
	strArr := make([]string, 4)

	appendedSlice := append(strArr)

	//it will slice from index 1 to 2 (last index will not be included
	slicedArr := appendedSlice[1:3]
	fmt.Println(slicedArr)

	//it will return slice containing elements less than index 4(4 index will be ignored)
	sliceArr1 := appendedSlice[:4]
	fmt.Println(sliceArr1)

	//We can declare and initialize a variable for slice in a single line as well.
	t := []string{"data", "res", "java"}
	fmt.Println(t)

	var t2 []string

	copy(t2, t)

	//slice package contain various methods that we can use on the slice
	if slices.Equal(t, t2) {
		println("t1 and t2 are equal")

	}

	//multi dimensional slice

	twoD := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(twoD)

	//slices can be composed of multidimensional data structure.
	// The length of internal slice can vary unlike in array

	//multi dimension slice
	twoD2 := make([][]int, 3)

	fmt.Println(twoD2)

	for i := 0; i < len(twoD2); i++ {

		innerLen := 3
		//creating internal slices and pushing them
		twoD2[i] = make([]int, innerLen)

		//pushing values in internal slice
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}

	}
	fmt.Println(twoD2)

}
