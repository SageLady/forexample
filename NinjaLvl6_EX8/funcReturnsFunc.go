package main

import "fmt"

func main()  {
y := bar(4)
fmt.Printf("The type: %T and the value %v", y, y(4))
i := y(4)
fmt.Println("\n",i)
fmt.Println(bar(4)(4))

}

func bar(x int) func(x int) int {
	return func(x int) int {
			var factorial = 1
			for x != 0 {
				factorial *= x
				x--
			} // END for
			return factorial
	} // END func
}  // END bar

//Hands-on exercise #8
//Create a func which returns a func
//assign the returned func to a variable
//call the returned func
//code: https://play.golang.org/p/c2HwqVE1Rd
//video: 109

//SOLUTION:
//package main

//import "fmt"

//func main() {
//	f := foo()
//	fmt.Println(f())
//}

//func foo() func() int {
//	return func() int {
//		return 42
//	}
//}
