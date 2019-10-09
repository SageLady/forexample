package main

import (
	"fmt"
)

//Channels can be
//1. Receive only
//2. Send only
//3. Both

//Read from left to right


func main() {
	c := make(chan int, 2) //bi-directional channel
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//General to Specific converts:
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))   //general to specific converts --conversion
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	//Note:
	//1. Bi-directional only assigns to bi-directional
	     //Example: c = d  (where d is bi-directional)
	     //e = make(<-chan int) receive channel
	     //c != e
	//2. Receive and send channels can only be assigned to receive or send channels
	//3. BUT - RECEIVER/SENDER can be assigned by bi-directional
	//	 Example: c = bi-directional
	//			  e = receive, and f = send
	//	 e = c
	//	 f = c

}


//CREATING A SEND CHANNEL
//func main() {
//	cs := make(chan<- int)

//	go func() {
//		cs <- 42
//	}()
//	fmt.Println(<-cs)

//	fmt.Printf("------\n")
//	fmt.Printf("cs\t%T\n", cs)
//}


//SEND AND RECEIVE one channel
//func main() {
//	c := make(chan int)

//	go func() {
//		c <- 42
//	}()

//	fmt.Println(<-c)
//}