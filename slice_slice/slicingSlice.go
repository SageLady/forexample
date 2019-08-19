package main

import "fmt"

func main()  {
	x := []int{0, 2, 3, 5, 6, 8, 10, 14, 17}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:4])

	for i := 0; i <= 8; i++ {
		fmt.Println(i, x[i])
	}

}
