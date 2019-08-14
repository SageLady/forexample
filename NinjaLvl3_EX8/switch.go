package main

import "fmt"

func main()  {
	switch {
	case 2 == 5:
		fmt.Println("2 does not equal 5, should not print")
	case 3 == 7 :
		fmt.Println("3 does not equal 7, should not print")
	case 4 == 4 :
		fmt.Println("4 does equal 4, this should print")
	} //END switch
} //END main

//Hands-on exercise #8
//Create a program that uses a switch statement with no switch expression specified.
//solution: https://play.golang.org/p/YpPgkWGqKY
//video: 057

//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {
//	switch {
//	case false:
//		fmt.Println("should not print")
//	case true:
//		fmt.Println("should print")
//	}
//}
