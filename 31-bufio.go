package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Normal io reads and writes one byte at a time but bufio writes when the buffer is filled. And all buffer is read at one or written in once
*/

func main(){

	f, err := os.Open("./text.txt")
	if err != nil {
		panic("FILE NOT FOUND")
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	fmt.Println(reader)
}
