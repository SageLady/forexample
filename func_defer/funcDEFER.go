package main

import "fmt"

func main()  {
	defer foo()  //When main closes then all defers get run
	bar()
}  //END main
func foo()  {
	fmt.Println("foo")
}
func bar()  {
	fmt.Println("bar")
}


//A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.