package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {  // write your error code here
		e := errors.New("more coffee needed")
			e = fmt.Errorf("more coffee needed - value was %v", f)
				return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
// e := errors.New("more coffee needed")
//		e := fmt.Errorf("more coffee needed - value was %v", f)
//		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
//	}
//	return 42, nil
//}


//Hands-on exercise #4
//Starting with this code, use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
//lat "50.2289 N"
//long "99.4656 W"
//solution:
//https://play.golang.org/p/nsRxbDfkCh
//video: 179

//SOLUTION:
//package main

//import (
//"fmt"
//"log"
//)

//type sqrtError struct {
//	lat  string
//	long string
//	err  error
//}

//func (se sqrtError) Error() string {
//	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
//}

//func main() {
//	_, err := sqrt(-10.23)
//	if err != nil {
//		log.Println(err)
//	}
//}

//func sqrt(f float64) (float64, error) {
//	if f < 0 {
		// e := errors.New("more coffee needed")
//		e := fmt.Errorf("more coffee needed - value was %v", f)
//		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
//	}
//	return 42, nil
//}


