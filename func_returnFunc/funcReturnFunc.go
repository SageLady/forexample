package main

import (
	"fmt"
)

func main()  {
	s1 := foo()   //Case: Returning a string
	fmt.Println(s1)

	x := func() int{
		return 51
	}
	fmt.Println(x)    //Value of x
	fmt.Printf("%T\n", x)  //Type x
	y := bar()
	fmt.Printf("%T\n", y)
	i := y()
	fmt.Println(i)

	fmt.Println(two())
	j := func() string{
		return "A string for j"
	}
	fmt.Println(j())
	fmt.Println(bar()()) //Print the value of the function that returns a function
}  //END main

func foo() string {   //Function returns a string
	s := "Hello World"
	return s
}
//type = func, identifier = bar() and return type: func() int
func bar() func() int {
	return func() int {
		return 51
	}
}

func two() string {
	return "A string to return"
}