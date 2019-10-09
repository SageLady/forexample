package main

//basic mutex - locks for both reading and writing
//RWmutex - more flexible allows for reading, but locks for writing
import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GO routines:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GO routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO routines:", runtime.NumGoroutine())
	fmt.Println("Counter value:", counter)
}  // END main



//Hands-on exercise #4
//Fix the race condition you created in the previous exercise by using a mutex
//it makes sense to remove runtime.Gosched()
//code: https://github.com/GoesToEleven/go-programming
//video: 151
//SOLUTION:
//package main

//import (
//"fmt"
//"sync"
//)

//func main() {
//	var wg sync.WaitGroup

//	incrementer := 0
//	gs := 100
//	wg.Add(gs)
//	var m sync.Mutex

//	for i := 0; i < gs; i++ {
//		go func() {
//			m.Lock()
//			v := incrementer
//			v++
//			incrementer = v
//			fmt.Println(incrementer)
//			m.Unlock()
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println("end value:", incrementer)
//}

