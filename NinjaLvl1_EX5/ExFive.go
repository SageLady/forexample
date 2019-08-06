package main

import "fmt"

type hotdog int
var x hotdog
var y int

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n",x) //This function does newline/cr
	x = 509              //Code puts the below next to above in print
	fmt.Println(x)  //This function does a carriage return/new line
	y = int(x)
	fmt.Println(y)  //new line
	fmt.Printf("%T",y) //new line


}

//Hands-on exercise #5
//Building on the code from the previous example
//1. at the package level scope, using the “var” keyword,
//create a VARIABLE with the IDENTIFIER “y”. The variable
//should be of the UNDERLYING TYPE of your custom TYPE “x”
//a. eg:
//type hotdog int

//var x hotdog
//var y int

//2. in func main
//a. this should already be done
//i.  print out the value of the variable “x”
//ii.  print out the type of the variable “x”
//iii.  assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//iv.  print out the value of the variable “x”
//b. now do this
//i.  now use CONVERSION to convert the TYPE of the
//VALUE stored in “x” to the UNDERLYING TYPE
//1. then use the “=” operator to ASSIGN that value to “y”
//2. print out the value stored in “y”
//3. print out the type of “y”
//code: here’s the solution: https://play.golang.org/p/cj8RrYgBOD
