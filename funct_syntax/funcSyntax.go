package main

import "fmt"

func main()  {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo()  {
	fmt.Println("hello from foo")
}
// Everything in go is PASS BY VALUE
func bar(s string)  {
	fmt.Println("hello,", s)
}

//RETURN
func woo(s string)string  {
		return fmt.Sprint("hello from woo, ", s)
}

//MULTI RETURN
func mouse(st string, ln string) (string, bool)  {
		a := fmt.Sprint(st, " ", ln, `says "Hello"`)
		b := true
		return a, b
}

// func ( r receiver) indentifier(parameters) (return(s)) {...}
// Parameters: When defining a function the parameters are
// defined. (if any exist)
//Arguments: When you call a function the parameters are now called
//arguments. We call our function and pass arguments.


