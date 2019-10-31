package main

import "fmt"

func main()  {
	c := make(chan int)

	// send
	//	go foo(c)
	//instead of using func foo going to put for loop here, so we can see scope better:
	go func() {
		for i := 0; i < 100; i++{
			c <- i
		}
		close(c)  //Without this close, receive is in hold mode waiting for more date to be
		//received. The for loop pulls until the channel is closed. It leaves the range loop
		//when the close is hit.
	}()

	// receive
	//bar(c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send
//func foo(c chan<- int)  {
//	for i := 0; i < 100; i++ {
//			c <- i
//	}
//	close(c)
//}

// receive -- this gets replaced by the for loop using range above
//func bar(c <-chan int)  {
//	fmt.Println(<-c)
//}
//Hands-on exercise #6
//write a program that
//puts 100 numbers to a channel
//pull the numbers off the channel and print them
//solution: https://play.golang.org/p/096Lk1BR7o 
//video: 169
