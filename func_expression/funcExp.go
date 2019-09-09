package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello")
	g := func(x int){
		fmt.Println("The beginning of the best", x)
	}
	g(51)
}


