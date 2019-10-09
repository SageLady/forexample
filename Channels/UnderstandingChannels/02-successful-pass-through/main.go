package main

import "fmt"

//Code DOES work
//Code works because there is the loading of the channel occuring inside the go function
//Which acts like a block
//The go func runs separately from the unloading of the channel, done by fmt.Println

func main()  {
	c := make(chan int)
	go func() {   //go function
		c <- 42   //Putting 42 on/in the channel. Here the channel blocks
	}()
	fmt.Println(<-c)  //Here again is some blocking where 42 is taken from channel
}