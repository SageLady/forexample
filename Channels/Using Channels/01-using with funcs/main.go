package main

import "fmt"

func main()  {
	c := make(chan int)
	// send
	go foo(c)   //goroutine function. Runs concurrently with main() thereby assigning value to channel
	// receive  //removing go from bar function because bar holds until it can do what it needs to do
	bar(c)   // This is a blocking function which I think waits for the goroutine to do it's job first.
	 fmt.Println("about to exit")
}

// send
func foo(c chan<- int)  {
	c <- 42
}

// receive
func bar(c <-chan int)  {
	fmt.Println(<-c)
}

//Using channels
//create general channels
//in funcs you can specify
//receive channel
//you can receive values from the channel
//a receive channel parameter
//in the func, you can only pull values from the channel
//you can’t close a receive channel
//send channel
//you can push values to the channel
//you can’t receive/pull/read from the channel
//you can only push values to the channel
//code:
//https://play.golang.org/p/t1rc8rSrMd
//code: https://github.com/GoesToEleven/go-programming
//video: 157

//Understanding channels
//Channels Introduction
//making a channel
//c := make(chan int)
//putting values on a channel
//c <- 42
//taking values off of a channel
//<-c
//buffered channels
//c := make(chan int, 4)
//channels  block
//they are like runners in a relay race
//they are synchronized
//they have to pass/receive the value at the same time
//just like runners in a relay race have to pass / receive the baton to each other at the same time
//one runner can’t pass the baton at one moment
//and then, later, have the other runner receive the baton
//the baton is passed/received by the runners at the same time
//the value is passed/received synchronously; at the same time
//channels allow us to pass values between goroutines
//code:
//doesn’t work
//https://play.golang.org/p/XPgsj2xS0F
//IMPORTANT: CHANNELS BLOCK
//channels allow
//coordination / synchronization / orchestration
//buffering (buffered channels)
//send & receive
//https://play.golang.org/p/SHr3lpX4so
//buffer
//https://play.golang.org/p/hsttb2qEJi
//“The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.” Golang Spec
//buffer doesn’t work
//https://play.golang.org/p/epLsvcivJS
//buffer
//https://play.golang.org/p/_6bSl5fc17
//code: https://github.com/GoesToEleven/go-programming
//video: 155


//Directional channels
//Channel type. Read from left to right.
//code:
//seeing the type
//	code from previous video
//https://play.golang.org/p/a98otBr4eX
//send & receive (bidirectional)
//https://play.golang.org/p/di1mKkGF6Y
//“send and receive” means “send and receive”
//https://play.golang.org/p/SHr3lpX4so
//already saw the above code
//send means send
//error: “invalid operation: <-cs (receive from send-only type chan<- int)”
//https://play.golang.org/p/oB-p3KMiH6
//receive means receive
//error: “invalid operation: cr <- 42 (send to receive-only type <-chan int)”
//https://play.golang.org/p/_DBRueImEq
//“A channel may be constrained only to send or only to receive [general to specific] by conversion or assignment.” Golang Spec
//doesn’t assign
//specific to general
//https://play.golang.org/p/bG7H6l03VQ
//specific to specific
//https://play.golang.org/p/8JkOnEi7-a
//assigns
//general to specific
//https://play.golang.org/p/hrNZq961KA
//conversion
//general to specific works
//https://play.golang.org/p/H1uk4YGMBB
//specific to general doesn’t work
//https://play.golang.org/p/4sOKuQRHq7
//code: https://github.com/GoesToEleven/go-programming
//video: 156

