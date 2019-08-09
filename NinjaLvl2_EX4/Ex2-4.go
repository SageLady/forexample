package main

import "fmt"

var x int

func main()  {
	x = 3  //==> x := 2 represents assigning a variable an int
	fmt.Printf("%d\t%b\t%#x", x, x, x)

	y:= x << 1
	fmt.Printf("\t%d\t%b", y, y)
}

//Hands-on exercise #4
//Write a program that
//1. assigns an int to a variable
//2. prints that int in decimal, binary, and hex
//3. shifts the bits of that int over 1 position to the left, and assigns that to a variable
//4. prints that variable in decimal, binary, and hex
//solution: https://play.golang.org/p/Ms964T8SbH
