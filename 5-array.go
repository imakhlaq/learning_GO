package main

import "fmt"

func main() {

	//an array is a numbered sequence of elements of a specific length

	var a [5]int //here we are created an array having 5 elements
	//TYPE of array and Length of array are part of array type
	// default values in side array is 0 in case of int

	fmt.Println(a)

	fmt.Println(a[2]) //accessing elements

	fmt.Println(len(a)) // length of an array

	// array with values ( initialising array with one line )
	agents := [4]string{"neon", "viper", "fade", "sage"}

	fmt.Println(agents)

	//another way initialising array with one line
	//compiler will figure out how many values does array have
	num := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(num)

	//u can set the specific index to
	//here we have skiped the index 2,3 it will have default value 0
	data := [...]int{1, 2, 4: 4, 5, 6}

	fmt.Println(data[2])

	//multi dimension array
	var twoD [2][2]int

	for i := 0; i < len(twoD[0]); i++ {

		for j := 0; j < len(twoD[0]); j++ {

			fmt.Println(twoD[i][j])
		}
	}

	//initializing 2d with values
	// u cant use [...]
	twoD2 := [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(twoD2)
}
