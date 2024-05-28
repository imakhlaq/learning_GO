package main

import "fmt"

// declaring an interface
type shape interface {
	clacArea() float64
	clacWidth() float64
}

// declaring some structs
type circle struct {
	area  float64
	width float64
}

type square struct {
	area  float64
	width float64
}

//to implement the interface u just have to add the methods of interface.
//in the struct

func (s *circle) clacArea() float64 {

	return s.width * s.area
}

func (s *circle) clacWidth() float64 {

	return s.width * s.area
}

func (s *square) clacArea() float64 {

	return s.width * s.area
}

func (s *square) clacWidth() float64 {
	return s.width * s.area
}

// function that calls mehods using interface
func mesure(g shape) {
	fmt.Println(g)
	fmt.Println(g.clacWidth())
	fmt.Println(g.clacArea())

}

func main() {

	c := circle{area: 1.145, width: 2.33}
	//circle and square is both implementing the shape interface
	//because in methods we are using pointer reference we have to pass pointer to the circle
	mesure(&c)

}
