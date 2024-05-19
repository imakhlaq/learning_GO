package main

import "fmt"

func main() {

	arr := []int{2, 3, 4}
	sum := 0

	//index,value
	for _, element := range arr { //skipping index by _
		sum += element
	}
	fmt.Println(sum)

	for index, element := range arr {

		if index == 2 {
			fmt.Println(element)
		}
	}

	//===========================================================================================================

	m := map[int]string{1: "kepla", 2: "sepla"}

	//using range ove a map
	for key, value := range m {

		fmt.Println(key, value)
	}
	// u can skip value , but when u have to skip key u have to use _ .
	// in place of key

	//iterating over string
	// it gives index and unicode
	for index, unicode := range "golang" {

		fmt.Println(index, unicode)
	}

}
