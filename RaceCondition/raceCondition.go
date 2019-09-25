package main
//This program is a representation of a
//RACE CONDITION using the variable counter

//If compiling from command line and you want to see
//if there is a race condition use flag race
//go run -race raceCondition.go
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GO routines:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GO routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO routines:", runtime.NumGoroutine())
	fmt.Println("Counter value:", counter)
}  // END main
