package main

import "fmt"

var x int     //These can be declared in main function using :=
var y float64
var z uint8  //u stands for unsigned uint8 goes from 0 - 255 (only these values can be assigned to this variable of this type.
//byte is an alias for uint8
var a int8 //int8  is for signed goes from -128 through 127, if 128 used it would not work
//Predeclared numeric types: uint = 32/64 bits; int = same size as uint; uintptr = an unsigned integer large enough to store the  uniterpreted bits of a pointer value

func main(){
	x = 2
	y = 2.34567
	z = 123
	a = -45
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", a)
}

