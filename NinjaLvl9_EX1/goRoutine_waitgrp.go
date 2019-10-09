package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main()  {
	wg.Add(1)
	go foo()
	bar2()
	foo2()
	go bar()
	wg.Wait()
}

func foo()  {
	for i := 0; i < 3; i++ {
		fmt.Println("Inside Foo:", i)
	}
	wg.Done()
} //END foo

func foo2()  {
	for i := 0; i < 3; i++ {
		fmt.Println("Inside Foo2:", i)
	}
} //END foo

func bar()  {
	for i:= 0; i < 2; i++ {
		fmt.Println("Inside Bar:", i)
	}
	wg.Done()
} //END bar


func bar2()  {
	for i:= 0; i < 2; i++ {
		fmt.Println("Inside Bar2:", i)
	}
} //END bar

//Hands-on exercise #1
//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists
//code: https://github.com/GoesToEleven/go-programming
//video: 148

//SOLUTION:
//package main

//import (
//"fmt"
//"runtime"
//"sync"
//)

//func main() {

//	fmt.Println("begin cpu", runtime.NumCPU())
//	fmt.Println("begin gs", runtime.NumGoroutine())

//	var wg sync.WaitGroup
//	wg.Add(2)

//	go func() {
//		fmt.Println("hello from thing one")
//		wg.Done()
//	}()

//	go func() {
//		fmt.Println("hello from thing two")
//		wg.Done()
//	}()

//	fmt.Println("mid cpu", runtime.NumCPU())
//	fmt.Println("mid gs", runtime.NumGoroutine())

//	wg.Wait()

//	fmt.Println("about to exit")
//	fmt.Println("end cpu", runtime.NumCPU())
//	fmt.Println("end gs", runtime.NumGoroutine())
//}