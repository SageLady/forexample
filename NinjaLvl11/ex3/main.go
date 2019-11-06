//Hands-on exercise #3
//Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
//HINT:
package main

import "fmt"

type customErr struct {
	detail string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.detail)
}

func main() {
	c1 := customErr{
		detail: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
}


//solution:
//ONE:
//https://play.golang.org/p/ixeowY2fd2
//TWO:
//assertion
//https://play.golang.org/p/pbl2kCYsM0
//THREE:
//conversion
//https://play.golang.org/p/1ldiBdkdzA
//video: 178


//Solution One:
//package main

//import (
//	"fmt"
//)

//type customErr struct {
//	info string
//}

//func (ce customErr) Error() string {
//	return fmt.Sprintf("here is the error: %v", ce.info)
//}

//func main() {
//	c1 := customErr{
//		info: "need more coffee",
//	}

//	foo(c1)
//}

//func foo(e error) {
//	fmt.Println("foo ran -", e, "\n")
//}

//SOLUTION TWO:
//package main

//import (
//	"fmt"
//)

//type customErr struct {
//	info string
//}

//func (ce customErr) Error() string {
//	return fmt.Sprintf("here is the error: %v", ce.info)
//}

//func main() {
//	c1 := customErr{
//		info: "need more coffee",
//	}

//	foo(c1)
//}

//func foo(e error) {
//	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
//}

//SOLUTION THREE:
//package main

//import (
//	"fmt"
//)

//type hotdog int

//func main() {
//	var x hotdog = 42
//	fmt.Println(x)
//	fmt.Printf("%T\n", x)

//	var y int
//	y = int(x)
//	fmt.Println(y)
//	fmt.Printf("%T", y)
//}

//	fmt.Println("foo ran -", e, "\n", e.(customErr).info)



