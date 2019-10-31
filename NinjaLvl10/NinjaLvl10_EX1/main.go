package main

import "fmt"

//func main() {
//	c := make(chan int)

//	c <- 42

//	fmt.Println(<-c)
//}

func main()  {
	c := make(chan int)
	// send
	go foo(c)
	bar(c)
	fmt.Println("about to exit")
}  // END main

// send
func foo(c chan<- int)  {
	c <- 42
}

// receive
func bar(c <-chan int)  {
	fmt.Println(<-c)
}


//Hands-on exercise #1
//get this code working:
//using func literal, aka, anonymous self-executing func
//solution: https://play.golang.org/p/SHr3lpX4so
//a buffered channel
//solution: https://play.golang.org/p/Y0Hx6IZc3U
//video: 164

//**using func literal
//package main

//import (
//"fmt"
//)

//func main() {
//	c := make(chan int)

//	go func() {
//		c <- 42
//	}()

//	fmt.Println(<-c)
//}

//**a buffered channel
//package main

//import (
//"fmt"
//)

//func main() {
//	c := make(chan int, 1)

//	c <- 42

//	fmt.Println(<-c)
//}