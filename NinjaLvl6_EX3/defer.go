package main

import "fmt"

func main()  {
fmt.Println("The proper order of greeting the day:")
defer foo()
bar()
}

func foo()  {
	defer func() {
				fmt.Println("\nFinally, Good-Bye and Good-Night")
			}()
	fmt.Println("\nNot ready to close out the day...")
}

func bar()  {
	fmt.Println("\nGood Morning Sunshine!")
}

//Hands-on exercise #3
//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
//code: https://play.golang.org/p/XluEuUD0Nw
//video: 104

//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {
//	defer foo()
//	fmt.Println("Hello, playground")
//}

//func foo() {
//	defer func() {
//		fmt.Println("foo DEFER ran")
//	}()
//	fmt.Println("foo ran")
//}

