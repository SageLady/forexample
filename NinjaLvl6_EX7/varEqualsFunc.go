package main

import "fmt"

func main()  {
f := func(x int) {
	fmt.Println("The best is yet to come...", x)
}  // END func
f(52)

} // END main


//Hands-on exercise #7
//Assign a func to a variable, then call that func
//code: https://play.golang.org/p/_Qu7ZAyFDH
//video: 108


//SOLUTION:
//package main

//import (
//"fmt"
//)

//var x int
//var g func()

//func main() {

//	f := func() {
//		for i := 0; i < 3; i++ {
//			fmt.Println(i)
//		}
//	}

//	f()
//	fmt.Printf("%T\n", f)

//	fmt.Println(x)
//	fmt.Printf("%T\n", x)

//	g = f
//	g()
//	fmt.Printf("this is g %T\n", g)
//
//	fmt.Println("done")
//}

