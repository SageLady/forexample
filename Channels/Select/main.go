package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel  (sending on to the channel)
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

// receive channel  (receiving onto the channel)
func receive(even, odd <-chan int, quit <-chan bool) {
	for {      // Infinite loop
		select {    //Pulling values of channels
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case <-quit:
			return
		}
	}
}




//Channels BLOCK
//RANGE clause also blocks and is used to range over a channel


//Select
//Select statements pull the value from whatever channel has a value ready to be pulled.
//code:
//building on previous code
//https://play.golang.org/p/41m5Mt7B-5
//not closing even odd channels
//https://play.golang.org/p/Emx6S4zn_y
//code: https://github.com/GoesToEleven/go-programming
//video: 159
