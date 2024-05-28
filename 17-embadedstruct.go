package main

import "fmt"

type base struct {
	number string
}

func (b *base) PrintName() {
	fmt.Println(b.number)

}

// container
type container struct {
	name string
	id   int

	//embadding
	base
}

func main() {

	// instantiation container
	c := container{
		id:   11,
		name: "akhlaq",
		base: base{number: "7905399065"},
	}

	//accessing base struct property in container struct
	c.PrintName()

	//changing it
	c.name = "huuu"

	//accessing base struct property in container struct
	//    c.number

	fmt.Printf("%+v", c)

}
