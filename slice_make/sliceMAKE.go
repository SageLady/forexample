package main

import "fmt"

func main()  {
	x := make([]int, 5, 15)  //[]T (type, length, capacity)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 34
	x[1] = 56
	fmt.Println(x)
	x[4] = 88
	fmt.Println(x)
	x = append(x, 12, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 14, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 16, 17)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

//The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function cap(a).
