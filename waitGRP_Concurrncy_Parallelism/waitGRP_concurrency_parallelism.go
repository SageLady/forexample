package main

import (
	"fmt"
	"runtime"    //runtime functions and vars
	"sync"       //go routine and wait groups
)
var wg sync.WaitGroup

func main()  {
		fmt.Println("OS\t\t", runtime.GOOS)
		fmt.Println("ARCH\t\t", runtime.GOARCH)   //Architecture
		fmt.Println("CPUs\t\t", runtime.NumCPU())   //Function - Number of CPUs
		fmt.Println("Go Routines\t", runtime.NumGoroutine())   //Function - Number of GO routines running
		 wg.Add(1)
		 go foo()   // This is now a go routine
		 bar()
		fmt.Println("CPUs\t\t", runtime.NumCPU())   //Function - Number of CPUs
		fmt.Println("Go Routines\t", runtime.NumGoroutine()) //# of GO Routines
		wg.Wait()
} //END main

func foo()  {
	for i := 0; i < 10; i++ {
			fmt.Println("foo:", i)
	}
	wg.Done()
} //END foo
func bar()  {
	for i := 0; i < 10; i++ {
			fmt.Println("bar:", i)
	}
} //END bar

//Functions to be used:
//number of CPUs
//number of GO routines