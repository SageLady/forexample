package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

//Fan in
//Taking values from many channels, and putting those values onto one channel.
//code:
//Todd’s code
//https://play.golang.org/p/_CyyXQBCHe
//Rob Pike’s code
//https://play.golang.org/p/buy30qw5MM
//video: 161
//Fan out
//Taking some work and putting the chunks of work onto many goroutines.
//code:
//fan out in
//https://play.golang.org/p/iU7Oee2nm7
//throttle throughput
//https://play.golang.org/p/RzR3Kjrx7q
//video:  162
//Context

//ROB PIKE'S CODE:
//package main

//import (
//"fmt"
//"math/rand"
//"time"
//)

//func main() {
//	c := fanIn(boring("Joe"), boring("Ann"))
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-c)
//	}
//	fmt.Println("You're both boring; I'm leaving.")
//}

//func boring(msg string) <-chan string {
//	c := make(chan string)
//	go func() {
//		for i := 0; ; i++ {
//			c <- fmt.Sprintf("%s %d", msg, i)
//			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//		}
//	}()
//	return c
//}

// FAN IN
//func fanIn(input1, input2 <-chan string) <-chan string {
//	c := make(chan string)
//	go func() {
//		for {
//			c <- <-input1
//		}
//	}()
//	go func() {
//		for {
//			c <- <-input2
//		}
//	}()
//	return c
//}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
