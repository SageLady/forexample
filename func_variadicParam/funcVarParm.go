package main

import "fmt"

func main()  {
	x := sum(2,3,4,5,6,7,8,9)
	fmt.Println("The total is, ", x)
}

//Stores as a slice of int
func sum(x ...int)int  {  //The type here is a slice of int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0   //initializing sum as zero
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i,
			"we are now adding", v, "to the total which is now",
			sum)
	} //END for

	return sum
} //END foo
//func (r receiver) identifier(params(s)) return(s) {}