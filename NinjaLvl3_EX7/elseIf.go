package main

import "fmt"

func main() {

	if 3 == 5 {
		fmt.Println("Not equal to 5")
	} else if 3 == 2 {
		fmt.Println("Not equal to 2")
	} else {
		fmt.Println("Three does not equal 5 or 2")
	}

} //END main

//Hands-on exercise #7
//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
//solution: https://play.golang.org/p/IDnrJpE7vT
//video: 056

//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {
//	x := "Moneypenny"

//	if x == "Moneypenny" {
//		fmt.Println(x)
//	} else if x == "James Bond" {
//		fmt.Println("BONDDONBONDONBOND", x)
//	} else {
//		fmt.Println("neither")
//	}
