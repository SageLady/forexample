//package errorsWithInfo

//import (
//	"errors"
//	"fmt"
//	"log"
//)

//Errors with info
//We can add information to our errors. We can do this with
//errors.New()
//fmt.Errorf()
//builtin.error
//“Error values in Go aren’t special, they are just values like any other, and so you have the entire language at your disposal.” - Rob Pike
//code:
//https://github.com/GoesToEleven/go-programming
//video: 175


//ERRORS NEW
//CODE EXAMPLE ONE:
//package main

//import (
//"errors"
//"log"
//)

//func main() {
//	_, err := sqrt(-10)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, errors.New("norgate math: square root of negative number")
//	}
//	return 42, nil
//}


//CODE EXAMPLE TWO:
//Errors new var
//Same as above, but seeing the type and with a variable
//package main

//import (
//"errors"
//"fmt"
//"log"
//)

//var ErrNorgateMath = errors.New("norgate math: square root of negative number")

//func main() {
//	fmt.Printf("%T\n", ErrNorgateMath)
//	_, err := sqrt(-10)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, ErrNorgateMath
//	}
//	return 42, nil
//}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go

//CODE EXAMPLE THREE
//fmt.Errorf ==> Error formats according to a format specifier and returns the string
//Calls error new
//String with format specifier, returns error
//Uses polymorphism
//package main

//import (
//"fmt"
//"log"
//)

//func main() {
//	_, err := sqrt(-10.23)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
//	}
//	return 42, nil
//}

//Code Example four
//Assigning to a variable. Same as above, but use of var//package main
//Method attached to a type--like a struct with the struct signature--then
//method implicitly implements that interface.
//fmt.Errorf being used


//import (
//"fmt"
//"log"
//)

//func main() {
//	_, err := sqrt(-10.23)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
//		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
//		return 0, ErrNorgateMath
//	}
//	return 42, nil
//}

//CODE EXAMPLE FIVE:
//package main

//import (
//"fmt"
//"log"
//)

//type norgateMathError struct {
//	lat  string
//	long string
//	err  error
//}

//func (n norgateMathError) Error() string {
//	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
//}

//func main() {
//	_, err := sqrt(-10.23)
//	if err != nil {
//		log.Println(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
//		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
//		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
//	}
//	return 42, nil
//}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go