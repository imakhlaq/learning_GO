package main

import (
	"fmt"
	"time"
)

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
	//switch

	i := 3

	switch i {
	case 1:
		fmt.Println("LOW")
	case 2:
		fmt.Println("2")
	case 4:
		fmt.Println("3")
	default:
		fmt.Println("DEFAULT")

	}

	day := time.Now().Weekday() //using data time getting the day of the week

	switch day {

	//checking which day it is
	case time.Sunday, time.Saturday:
		fmt.Println("WEEKEND")

	default:
		fmt.Println("WEEKDAY")

	}

	//you can use switch to run a code base on the condition (if-else replacement)
	t := time.Now()
	//====||                      no point to switch
	switch {

	case t.Hour() < 12:
		fmt.Println("noon")

	case t.Hour() > 12:
		fmt.Println("evening")
	}

}
