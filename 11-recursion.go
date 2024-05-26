package main

import (
	"fmt"
)

func calcFib(num int) int64 {

	if num == 0 {
		return 0
	}
	if num == 1 || num == 2 {
		return 1
	}

	return calcFib(num-1) + calcFib(num-2)
}

func main() {

	fibOfNum := calcFib(30)

	fmt.Println(fibOfNum)
}
