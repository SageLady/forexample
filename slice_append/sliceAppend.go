package main

import "fmt"

func main() {
	x := []int{0, 2, 3, 5}
	fmt.Println(x)
	x = append(x, 99, 34, 45)
	fmt.Println(x)

	y := []int{34, 22, 11, 55, 66}
	x = append(x, y...)   //variadic parameter - unlimited values of type int or whatever type.
	fmt.Println(x)
}
	//Effective go:
	//What append does is append the elements to the end of the slice and return the result. The result needs to be returned because, as with our hand-written Append, the underlying array may change. This simple example


