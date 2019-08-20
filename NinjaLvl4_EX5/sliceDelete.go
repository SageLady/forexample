package main

import "fmt"

func main()  {//0   1  2   3   4   5   6   7   8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[0:3], x[6:]...)
	fmt.Println(y)
}

//Hands-on exercise #5
//To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
//start with this slice
//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
//[42, 43, 44, 48, 49, 50, 51]
//solution: https://play.golang.org/p/u8zpHLfgSE
//just for fun:
//https://goo.gl/frz786
//https://www.youtube.com/channel/UCJ8YIwWQCO7hMiqpOw2ZLFw
//video: 075
