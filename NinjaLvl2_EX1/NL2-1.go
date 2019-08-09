package main

import "fmt"

func main() {
x := 42
fmt.Printf("%b\t", x)   //binary
fmt.Printf("%#X\t", x)    //hexadecimal
fmt.Printf("%d", x)       //decimal
}

//PROBLEM:
// Hands On exercise #1:
//Write a program that prints a number in decimal, binary
//and hex.
//solutions: solution: https://play.golang.org/p/bAQxcEuK8O

//Solution:
//package main

//import (
//"fmt"
//)

//func main() {
//	a := 42
//	fmt.Printf("%d\t%b\t%#x", a, a, a)
//}

//Search for solution: Found what I needed here
//https://golang.org/doc/effective_go.html#type_switch
//https://golang.org/doc/effective_go.html#formatting
//https://golang.org/doc/effective_go.html#printing
//var t interface{}
//t = functionOfSomeType()
//switch t := t.(type) {
//default:
//fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
//case bool:
//fmt.Printf("boolean %t\n", t)             // t has type bool
//case int:
//fmt.Printf("integer %d\n", t)             // t has type int
//case *bool:
//fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
//case *int:
//fmt.Printf("pointer to integer %d\n", *t) // t has type *int
//}
