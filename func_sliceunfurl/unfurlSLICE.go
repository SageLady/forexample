package main

import "fmt"
//The final incoming parameter in a function signature may have a type prefixed with .... A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.
//variadic params must be the final parameter
func main()  {
	x := sum("string") //dots allow the slice of int to be used as int arg
	fmt.Println("The total is, ", x)

}

//Stores as a slice of int
func sum(s string, x ...int)int  {  //The type here is a slice of int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0   //initializing sum as zero
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i,
			"we are now adding", v, "to the total which is now",
			sum)
	} //END for
	return sum
} //END foo