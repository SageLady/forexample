package main

import "fmt"

func main()  {
	fmt.Println(4*3*2*1)
//	n := factorial(4)
//	fmt.Println(n)
	v := factLoop(4)
	fmt.Println("\n")
	fmt.Println(v)


}  // END main

func factorial(n int) int {
	if n == 0 {
		return 1
	} // END if
	return n * factorial(n-1)
}  // END factorial

func factLoop(n int) int {
	var factorial = 1
	for n != 0 {
		factorial *= n
		fmt.Printf("\nn equals %v", n)
		fmt.Printf("\nfactorial equals  %v", factorial)
		n--
	}
	return factorial
}

// 4 * 3 * 2 * 1 * 1

//Solution in class:
//func loopFact(n int) int {
//	total := 1
//	for ; n > 0; n-- {
//		total *= n
//	}
//	return total
//}
