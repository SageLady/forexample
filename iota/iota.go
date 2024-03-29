package main

import "fmt"
//One way to initialize:
//const (
//	a = iota
//	b = iota
//	c = iota
//)

//Another way to initialize:
const (
	a = iota
	b
	c
)
const (
	d = iota
	e
	f
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}


//predeclared identifier
