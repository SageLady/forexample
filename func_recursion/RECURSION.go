package main

import "fmt"

func main()  {
	fmt.Println(4*3*2*1)
	n := factorial(4)
	fmt.Println(n)


}  // END main

func factorial(n int) int {
	if n == 0 {
		return 1
	} // END if
	return n * factorial(n-1)
}  // END factorial

func factLoop(n int) int {
	for i == 0; i <= n; i++ {
		

	}

}

// 4 * 3 * 2 * 1 * 1
