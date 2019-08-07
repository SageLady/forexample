package main

import "fmt"

//bit shifting is to take 0000 0010 binary number
//and                     0000 0100 shift the one
//bit shifting operators: << and >>

func main()  {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y:= x << 1
    fmt.Printf("%d\t\t%b", y, y)
}

