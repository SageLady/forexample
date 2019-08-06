package main

import "fmt"
var x int = 42
var y string = "James Bond"
var z bool = true

func main()  {
	s := fmt.Sprintf("%v\t%v\t%v",x,y,z)
	fmt.Println(s)
}

//%v = for value, not a conversion of type

//Hands-on exercise #3
//Using the code from the previous exercise,
//1. At the package level scope, assign the following values to the three variables
//a. for x assign 42
//b. for y assign “James Bond”
//c. for z assign true
//2. in func main
//a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
//b. print out the value stored by variable “s”//code: here’s the solution: https://play.golang.org/p/QFctSQB_h3