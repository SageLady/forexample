package main

import "fmt"

const a  = 42
const b = 42.78
const c = "James Bond"

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
//Second way to declare constants:
// const (
//			a int = 42
//			b float = 42.78
//          c atring = "James Bond"
//  )

//Above is a typed constant which does not allow
//compiler to choose type, but developer does
