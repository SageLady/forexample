package main

import "fmt"

func main()  {
	x := []int{0, 2, 3, 5}
	fmt.Println(x)
	x = append(x, 99, 34, 45)
	fmt.Println(x)
	
	y := []int{23, 55, 77, 88, 99}
	x = append(x, y...)
	fmt.Println(x)
	//Grab from beginning of the slice
	//represented by : and then on
	//the other side add the location
	//in this exercise up to location 2 [:2]
	//we have our first slice x[:2] and then append to 
	//this slice position 7 forward 
	
	x = append(x[:2], x[7:]...)
	fmt.Println(x)
}
