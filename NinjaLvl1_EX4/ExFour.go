package main

import "fmt"

type hotdog int
var x hotdog

func main()  {
	fmt.Println(x)   //zero value
	fmt.Printf("%T\n",x)  //Printf() %T = type of variable x
	x = 42   //equal operator
    fmt.Println(x) //value stored in variable
}
//printf = use to print a "type of variable" which this is %T
//UNDERLYING TYPE - see link below, is terminology used to
//research creating your own type, as in this example

//Hands-on exercise #4
//FYI - nice documentation and new terminology “underlying type”
//https://golang.org/ref/spec#Types
//For this exercise
//1. Create your own type. Have the underlying type be an int.
//2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using
//the “VAR” keyword
//3. in func main
//a. print out the value of the variable “x”
//b. print out the type of the variable “x”
//c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
//d. print out the value of the variable “x”
//code: here’s the solution: https://play.golang.org/p/snm4WuuYmG