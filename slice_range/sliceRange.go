package main

import "fmt"

func main()  {
	x := []int{0, 2, 3, 5, 6, 8, 10, 14, 17}
	fmt.Println(len(x))  //length of slice
	fmt.Println(cap(x))
	fmt.Println(x)
	fmt.Println(x[5])   //accessing by index position
for i, v := range x {
	fmt.Println(i,v)
}

}
