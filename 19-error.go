package main

import "fmt"

// defining custom err
type CustomErr struct {
	statusCode int
	message    string
}

// implementing Error interface
func (c CustomErr) Error() string {
	return c.message
}

// ========================|| error interface
func calc(num int) (int, error) { //go convention is to return second value as error

	if num >= 44 {
		//convention
		return -1, CustomErr{statusCode: 301, message: "Number is small"}
	}
	return num, nil
}

func main() {

	res, err := calc(100)
	//checking error
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)

}
