package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)   //Can be an int: quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel
//If quit is an int: func send(even, odd, quit chan<- int)
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
//If quit is an int: func receive(even, odd, quit <-chan int) {
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}


// Simple comma OK idiom program:
// The comma ok idiom is used for checking status of a channel.

//package main

//import (
//"fmt"
//)

//func main() {
//	c := make(chan int)
//	go func() {
//		c <- 42
//	}()

//	v, ok := <-c

//	fmt.Println(v, ok)
//}


// Simple comma ok idiom program
// Uses channels.
// Uses close
//package main

//import (
//"fmt"
//)

//func main() {
//	c := make(chan int)
//	go func() {
//		c <- 42
//		close(c)
//	}()

//	v, ok := <-c

//	fmt.Println(v, ok)

//	v, ok = <-c

//	fmt.Println(v, ok)
//}


//Code below uses quit as an int.
//Comma ok appears to be a bool as the output from
//the code below is false when ok is print upon termination of code
//However, it looks at the value of quit when 0 (zero) and acknowledges
//it as being false. When 1 it is true.

//package main

//import (
//"fmt"
//)

//func main() {
//	even := make(chan int)
//	odd := make(chan int)
//	quit := make(chan int)

//	go send(even, odd, quit)

//	receive(even, odd, quit)

//	fmt.Println("about to exit")
//}

// send channel
//func send(even, odd, quit chan<- int) {
//	for i := 0; i < 10; i++ {
//		if i%2 == 0 {
//			even <- i
//		} else {
//			odd <- i
//		}
//	}
//	close(quit)
//}

// receive channel
//func receive(even, odd, quit <-chan int) {
//	for {
//		select {
//		case v := <-even:
//			fmt.Println("the value received from the even channel:", v)
//		case v := <-odd:
//			fmt.Println("the value received from the odd channel:", v)
//		case i, ok := <-quit:
//			if !ok {
//				fmt.Println("from comma ok", i, ok)
//				return
//			} else {
//				fmt.Println("from comma ok", i)
//			}
//		}
//	}
//}


// quit is a bool
// infinite for loop and select used
// close used
// receive and send are functions
// send is go function
// receive is blocking function
// select statements used in each function

//package main

//import (
//"fmt"
//)

//func main() {
//	even := make(chan int)
//	odd := make(chan int)
//	quit := make(chan bool)

//	go send(even, odd, quit)

//	receive(even, odd, quit)

//	fmt.Println("about to exit")
//}

// send channel
//func send(even, odd chan<- int, quit chan<- bool) {
//	for i := 0; i < 100; i++ {
//		if i%2 == 0 {
//			even <- i
//		} else {
//			odd <- i
//		}
//	}
//	close(quit)
//}

// receive channel
//func receive(even, odd <-chan int, quit <-chan bool) {
//	for {
//		select {
//		case v := <-even:
//			fmt.Println("the value received from the even channel:", v)
//		case v := <-odd:
//			fmt.Println("the value received from the odd channel:", v)
//		case _, _ = <-quit:
//			return
//		}
//	}
//}



//Comma ok idiom
//The comma ok idiom with select.
//code:
//closing quit channel & comma ok idiom
//https://play.golang.org/p/fomc4Sc-gz
//with bool
//https://play.golang.org/p/QAwBLDKZuq
//with int
//https://play.golang.org/p/6XaTWxKpU3
//just comma ok idiom
//https://play.golang.org/p/6LPzCtZeT3
//https://play.golang.org/p/dToDc0zJhZ
//clean up above code - comma ok idiom
//step 1 - comma ok idiom code reduced
//https://play.golang.org/p/60K5-7xat6
//step 2 - remove underscore variable throwaways
//https://play.golang.org/p/QWzGUISkJK
//step 3 - change quit from bool to int
//https://play.golang.org/p/tFlvGg9ENT
//select receive
//https://play.golang.org/p/62vRH3jeQ6
//select send
//https://play.golang.org/p/bv_vCcJ6zs
//interesting
//doesn’t run
//https://play.golang.org/p/o7Quy6wc6r
//runs
//https://play.golang.org/p/8ZMDPBqHwt
//runs a different way - I like this one better
//quit channel was removed
//select was removed
//range over channel used instead
//https://play.golang.org/p/_CyyXQBCHe
//video: 160
