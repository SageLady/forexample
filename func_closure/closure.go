package main

import (
	"fmt"
)

var x int  // Scope of x is the entire code

func main()  {
	var y int    // Scope of y is main
	fmt.Println(y)   // y = 0
	fmt.Println(x)   // x = 0
	x++
	fmt.Println(x)   // x = 1
	foo()
	fmt.Println(x)   // x = 2

	{   // Code block in code block
		z := 51     // scope of z is braces, cannot use z outside of braces
		fmt.Println(z)
	}
	a := incrementor()
	b := incrementor()
	fmt.Println("Start a:")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("Start b:")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}  //END main

func foo() {
	x++
}

func incrementor() func() int  {
	 var p int // This code is only hit the first call
	 fmt.Printf("Inside incrementor %v\n", p) // This code is only hit the first call.
	 return func() int {   //First call and forward hit this code.
		 p++
		 return p
	 }  // END func
} // END incrementor
