package main

import "fmt"
//This is supposed to be like a while loop
func main()  {
	i := 1967
	for i < 2020 {
		fmt.Println(i)
		i++
	} //END for
} //END main

//Hands-on exercise #3
//Create a for loop using this syntax
//for condition { }
//Have it print out the years you have been alive.
//solution: https://play.golang.org/p/tnyqBPJ-i5
//video: 052

//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {
//	bd := 1985
//	for bd <= 2017 {
//		fmt.Println(bd)
//		bd++
//	}
//}
