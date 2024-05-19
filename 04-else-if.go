package main

import "fmt"

func main() {

	//if-else does not require brackets but braces is required
	k := 9
	if k == 8 || k > 10 {

		fmt.Println("if case")
	} else {
		fmt.Println("if case")
	}

	//nested else if
	if 9/2 == 0 {
		fmt.Println("ROUNDED")

	} else if 10 < 20 && 90 > 3 || 30 > 7 {
		fmt.Println("ELSE IF")
	} else if 20%1 == 0 {
		fmt.Println("ANOTHER ELSE IF")
	} else {
		fmt.Println("RANDOM")
	}

}
