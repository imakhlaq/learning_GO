package main

import "fmt"

func main() {

	//normal loop
	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}

	//range
	for i := range 5 {
		fmt.Println(i)
	}

	//range
	// you can also do this
	l := 4
	for k := range 8 {

		if k%2 == 0 {
			fmt.Println(k + l)
		}
	}

	//infinite
	for {
		fmt.Println("loop")
		break
	}
}
