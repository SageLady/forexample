package main

import "fmt"
//Slice declaration/ or initialization
func main()  {
	s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("The type is %T\n", s)
}
//Hands-on exercise #2
//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//Range over the slice and print the values out.
//Using format printing
//print out the TYPE of the slice
//solution: https://play.golang.org/p/sAQeFB7DIs
//video: 072

