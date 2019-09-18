package main

import "fmt"

func main()  {
s := multiply(3,6)
fmt.Println(s)
s2 := product(multiply,7,6)
fmt.Println("Product is:",s2)


}
//First arg: function, second int, third int
//Returns an int
func product(f func(x int, y int) int, x int, y int) int {
	return f(x,y)
}

//Two args are both ints
//Returns an int
func multiply(xi int, yi int) int {
		product := xi * yi
		return product
}

//Hands-on exercise #9
//A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument
//code: https://play.golang.org/p/0yGYPKh1y7
//video: 110


//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {

//	g := func(xi []int) int {
//		if len(xi) == 0 {
//			return 0
//		}
//		if len(xi) == 1 {
//			return xi[0]
//		}
//		return xi[0] + xi[len(xi)-1]
//	}

//	x := foo(g, []int{1, 2, 3, 4, 5, 6})
//	fmt.Println(x)
//}

//func foo(f func(xi []int) int, ii []int) int {
//	n := f(ii)
//	n++
//	return n
//}

