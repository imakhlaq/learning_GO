package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	//GO string literals are UTF-8 encoded text
	//string package have lots of functions for strings

	//in go strings are readonly slice of bites (means u cant change it)
	str1 := "string"
	fmt.Println(str1)

	var str2 string
	str2 = "dadad"
	fmt.Println(str2)

	//if you want to create a mutable string
	var str strings.Builder
	str.WriteString("dadadad")

	//in go character is known as rune
	var ru rune = 'a'
	fmt.Println(ru)

	//string is []bytes so u can do
	k := len(str1)
	fmt.Println(k)

	for i := 0; i < k; i++ {
		fmt.Println(str1[i]) // it will print the raw byte value of each string
	}

	//utf8 package provides usefully functions for string
	fmt.Println("unique character in string ", utf8.RuneCountInString(str1))

}
