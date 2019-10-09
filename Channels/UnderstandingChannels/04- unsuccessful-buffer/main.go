package main

import "fmt"

//Code DOES NOT work
//To fix code need buffer size of 2
//and two fmt.Println(<-c) statements
//LOADING = c <-42. make created channel with a buffer of size 1. So, that now we do not need the func to add value to the channel or create a storage for the int.

func main()  {
	c := make(chan int, 1)   //Buffered  channel allows a value to sit in the channel. Define size here
	c <- 42
	c <- 43  //Blocked because  only room for one value
	fmt.Println(<-c)  //Here again is some blocking where 42 is taken from channel

	//If you want to pull off 43 need another fmt.Println(<-c)
}
