package main

import "fmt"

//Code does not work
//It does not work because the channel needs to be able to
//load and unload at the same time
//If this cannot happen then the channel gets locked

func main()  {
 	c := make(chan int)
 	c <- 42
 	fmt.Println(<-c)
}
