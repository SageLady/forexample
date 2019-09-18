package main

import "fmt"

func main()  {
	foo()
	func() {
			fmt.Println("No args Anonymous says 'Wonderful Wednesday'")
	}()    // END func no args

	func(x int) {
			fmt.Println("I have an arg, the meaning of life is ", x)
	}(51)   // END func with arg
} // END main

func foo()  {
		fmt.Println("Inside of foo")
}
//Hands-on exercise #6
//Build and use an anonymous func
//code: https://play.golang.org/p/DQX3xEIcRe
//video: 107

//SOLUTION:

//package main

//import (
//"fmt"
//)

//func main() {

//	func() {
//		for i := 0; i < 100; i++ {
//			fmt.Println(i)
//		}
//	}()

//	fmt.Println("done")
//}