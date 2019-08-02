package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most" +
		"interesting class")
	foo()

	fmt.Println("something a little more here")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
} //main brace

func foo() {
	fmt.Println("I am in the foo function")
}

func bar() {
	fmt.Println("and then we exited")
}

//Control Flow
//(1) Sequence
//(2) loop; iterative
//(3) Conditional
