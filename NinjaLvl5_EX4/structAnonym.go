package main

import (
	"fmt"
)

func main()  {
	t1 := struct {
		name string
		age int
		gender string
	}{
		name: "Lucy",
		age: 2,
		gender: "male",
	}

	fmt.Println(t1)
	fmt.Printf("Gender is: %v", t1.gender)

} //END main

//SOLUTION:
//Hands-on exercise #4
//Create and use an anonymous struct.
//solution: https://play.golang.org/p/otBHFs-lPp
//video: 089

//package main

//import (
//"fmt"
//)

//func main() {

//	s := struct {
//		first     string
//		friends   map[string]int
//		favDrinks []string
//	}{
//		first: "James",
//		friends: map[string]int{
//			"Moneypenny": 555,
//			"Q":          777,
//			"M":          888,
//		},
//		favDrinks: []string{
//			"Martini",
//			"Water",
//		},
//	}
///	fmt.Println(s.first)
//	fmt.Println(s.friends)
//	fmt.Println(s.favDrinks)

//	for k, v := range s.friends {
//		fmt.Println(k, v)
//	}

//	for i, val := range s.favDrinks {
//		fmt.Println(i, val)
//	}
//}
