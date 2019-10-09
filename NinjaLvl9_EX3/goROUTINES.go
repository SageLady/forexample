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

//Hands-on exercise #3
//Using goroutines, create an incrementer program
//have a variable to hold the incrementer value
//launch a bunch of goroutines
//each goroutine should
//read the incrementer value
//store it in a new variable
//yield the processor with runtime.Gosched()
//increment the new variable
//write the value in the new variable back to the incrementer variable
//use waitgroups to wait for all of your goroutines to finish
//the above will create a race condition.
//Prove that it is a race condition by using the -race flag
//if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
//code: https://github.com/GoesToEleven/go-programming
//video: 150

//SOLUTION:
//package main

//import (
//"fmt"
//"runtime"
//"sync"
//)

//func main() {
//	var wg sync.WaitGroup

//	incrementer := 0
//	gs := 100
//	wg.Add(gs)

//	for i := 0; i < gs; i++ {
//		go func() {
//			v := incrementer
//			runtime.Gosched()
//			v++
//			incrementer = v
//			fmt.Println(incrementer)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println("end value:", incrementer)
//}
