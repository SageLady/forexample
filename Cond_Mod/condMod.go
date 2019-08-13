package main

import "fmt"

//loop, conditional, modulus or modulo
//example of getting even numbers using modulo operator
func main()  {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {   //any number divided by 2 with remainder 0
			fmt.Println(i)
		}
	}
}
