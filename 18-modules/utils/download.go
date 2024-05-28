package utils

import (
	"18-modules/types"
	"fmt"
)

func Download() {

	//u have to import it because its in different package
	u := types.User{Name: "akhlaq", Age: 3}

	// u cant import it because its private and only available in that package
	//   types.password

	fmt.Printf("%+v", u)

}
