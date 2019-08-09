package main

import "fmt"

const x int  = 12
const v bool = false

func main()  {
fmt.Printf("%v\t%v", x, v)
}

//Problem:
//Hands on #3:
//Create TYPED and UNTYPED constants.
//Print the values of the constants

//Solution: https://play.golang.org/p/NutvJXWUx2
//Solution: TYPED and UNTYPED
//package main

//import (
//"fmt"
//)

//const (
//	a = 42
//	b int = 43
//)

//func main() {
//	fmt.Println(a, b)
//}

//Definition of TYPED and UNTYPED constants
//Constants may be typed or untyped. Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped.

//A constant may be given a type explicitly by a constant declaration or conversion, or implicitly when used in a variable declaration or an assignment or as an operand in an expression. It is an error if the constant value cannot be represented as a value of the respective type.

//An untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a short variable declaration such as i := 0 where there is no explicit type. The default type of an untyped constant is bool, rune, int, float64, complex128 or string respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.
