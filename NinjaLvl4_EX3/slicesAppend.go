package main

import "fmt"

func main()  {//0   1   2   3   4   5   6   7   8   9
	s := []int{10, 11, 12, 42, 43, 44, 45, 46, 47, 48,
	//	10  11  12  13  14
		49, 50, 51, 18, 19}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("The type is %T\n", s)
	fmt.Println("Output should be: 42 43 44 45 46", s[3:8])
	fmt.Println("Output should be: 47 48 49 50 51", s[8:13])
	fmt.Println("Output should be: 44 45 46 47 48", s[5:10])
	fmt.Println("Output should be: 43 44 45 46 47", s[4:9])
	fmt.Println(s)
}

//Hands-on exercise #3
//Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//[42 43 44 45 46]
//[47 48 49 50 51]
//[44 45 46 47 48]
//[43 44 45 46 47]
//solution: https://play.golang.org/p/SGfiULXzAB
//video: 073