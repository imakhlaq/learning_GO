package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating a empty map
	//==========|| key[type]value Type
	m := make(map[string]int)

	//var name map[string]int

	//setting key and value
	m["key1"] = 22
	m["key2"] = 3

	//getting value of a key
	value := m["key1"]

	fmt.Println(value)

	//if you try to access a key that does not exist.
	// the default value of that map value type is returned. (i.e int 0)
	noExit := m["34"]
	fmt.Println(noExit)

	//no of key value pair in a map
	total := len(m)
	fmt.Println(total)

	//to delete the key and value from a map
	delete(m, "key2")

	// to remove all the key value pair in the map( i.e clear map)
	clear(m)

	//the optional second return value when u try to access the map key.
	//tells you that if key is present or not

	_, ok := m["Key2"] //leaving the value using _. use ok to check its available
	fmt.Println("is present", ok)

	//===========================================================================================================

	// You can declare and initialize a map in one line
	mk := map[int]string{1: "go", 2: "node"}

	fmt.Println(mk)

	ml := map[int]string{2: "cha", 3: "huuu"}
	fmt.Println(ml)

	//map package provide some utils functions
	maps.Equal(mk, ml)

}
