package types // Package types ==>types.go and dataReq.go are in same package.
import "fmt"

//================so u can use private stuff from types.go here and vice verse

// User is also defined in types.go and they both are in same scope, so they are in same scope.
// so they are in same name space. hence their will be name collision.
/*
type User struct {

}
*/

func creteUser() {

	//u don't need to import it because they are in same package
	p := password{pass: "12345", hash: "h25s"} //accessing private

	fmt.Printf("%+v", p)

}
