package main
//Solution one:
import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}

//Hands-on exercise #7
//write a program that
//launches 10 goroutines
//each goroutine adds 10 numbers to a channel
//pull the numbers off the channel and print them
//solutions:
//https://play.golang.org/p/R-zqsKS03P
//https://play.golang.org/p/quWnlwzs2z
//video: 170

//Solution 2:
//package main

//import (
//"fmt"
//"runtime"
//)

//func main() {
//	x := 10
//	y := 10
//	c := gen(x, y)

//	for i := 0; i < x*y; i++ {
//		fmt.Println(i, <-c)
//	}
//	fmt.Println("ROUTINES",runtime.NumGoroutine())
//}

//func gen(x, y int) <-chan int {
//	c := make(chan int)

//	for i := 0; i < x; i++ {
//		go func() {
//			for j := 0; j < y; j++ {
//				c <- j
//			}
//		}()
//		fmt.Println("ROUTINES",runtime.NumGoroutine())
//	}
//	return c
//}
