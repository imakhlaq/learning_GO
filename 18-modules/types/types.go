package types

import "fmt"

// User is public
type User struct {

	//name and age is public
	id   string
	Name string
	Age  int
}

// password is private
type password struct {

	//and the fields are also private
	pass string
	hash string
}

// PrintUser public func
func PrintUser(data User) {

	fmt.Printf("%+v", data)

}

// private func
func createPass() string {

	return "encrepted pass"

}
