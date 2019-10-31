package Recover

import "fmt"

//Notes:
//https://blog.golang.org/defer-panic-and-recover
//Go has the usual mechanisms for control flow: if, for, switch, goto. It also has the go statement to run code in a separate goroutine. Here I'd like to discuss some of the less common ones: defer, panic, and recover.

//A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

//CODE SAMPLE ONE:
//package main

//import "fmt"

//func main() {
//	var x int
//	x++
//	fmt.Println(x)
//	i := c()
//	fmt.Println(i)
//}

//func c() (i int) {
//	defer func() { i++ }()
//	return 1
//}

//CODE SAMPLE TWO:
//package main

//import "fmt"

//func main() {
//	f()
//	fmt.Println("Returned normally from f.")
//}

//func f() {
//	defer func() {
///		if r := recover(); r != nil {
	//		fmt.Println("Recovered in f", r)
	//	}
	//}()
	//fmt.Println("Calling g.")
	//g(0)
	//fmt.Println("Returned normally from g.")
//}

//func g(i int) {
///	if i > 3 {
	//	fmt.Println("Panicking!")
	//	panic(fmt.Sprintf("%v", i))
	//}
	//defer fmt.Println("Defer in g", i)
	//fmt.Println("Printing in g", i)
	//g(i + 1)
//}
