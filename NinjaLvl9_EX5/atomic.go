package main
//sync package contains atomic
//atomic allows us to avoid race conditions

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main()  {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GO routines:", runtime.NumGoroutine())
	var counter int64   //Anytime this type is present, think package atomic
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	//REPLACED BY PACKAGE ATOMIC ==>var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//REPLACED BY PACKAGE ATOMIC ==>mu.Lock()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//REPLACED BY PACKAGE ATOMIC ==>mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GO routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO routines:", runtime.NumGoroutine())
	fmt.Println("Counter value:", counter)
}  // END main

//Hands-on exercise #5
//Fix the race condition you created in exercise #4 by using package atomic
//code: https://github.com/GoesToEleven/go-programming
//video: 152

//SOLUTION:
//package main

//import (
//"fmt"
//"sync"
//"sync/atomic"
//)

//func main() {
//	var wg sync.WaitGroup
//	var incrementer int64

//	gs := 100
//	wg.Add(gs)

//	for i := 0; i < gs; i++ {
//		go func() {
//			atomic.AddInt64(&incrementer, 1)
//			r := atomic.LoadInt64(&incrementer)
//			fmt.Println(r)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println("end value:", incrementer)
//}


