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