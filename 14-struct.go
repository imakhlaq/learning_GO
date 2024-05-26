package main

import "fmt"

// struct are typed collection of fields. they are useful for grouping data

// this struct name and age fields
type person struct {
	name string
	age  int8
}

// this new person constructor create a new person and return its pointer
func newPerson(name string, age int8) *person {

	p := person{name: name, age: age}

	//here we are returning a pointer of a local variable.
	//You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	return &p
}

func main() {

	//create new struct
	fmt.Println(person{
		name: "akhlaq",
		age:  101,
	})

	//u can also omit the filed
	fmt.Println(person{"akhlaq", 35})

	//u can omit the value, but it will be 0 in case of int
	fmt.Println(person{name: "kalia"})

	//&prefix before gives you pointer to struct
	fmt.Println(&person{"superman", 40})

	//It's idiomatic to encapsulate new struct creation in constructor function (RECOMMENDED)
	fmt.Println(newPerson("akhlaq", 44))

	//struct are mutable and even if u have pointer of struct u can use . to access its properties
	s := person{name: "hula hu", age: 56}

	//storing pointer
	s1 := &s

	fmt.Println(s1.name)

	//changing in s1, which contain pointer to s
	s1.name = "akhlaq"

	fmt.Println(s.name)

	// if u have a struct that is used for one time u can skip its name
	k := struct {
		name string
		age  int8
	}{
		"kalia",
		22,
	}

	fmt.Println(k)

}
