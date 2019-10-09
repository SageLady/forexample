package main
//code: https://play.golang.org/p/Lmbyn7bO7e
import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}



//Context
//In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request's deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using. At Google, we developed a context package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as context. This article describes how to use the package and provides a complete working example.
//further reading:
//https://blog.golang.org/context
//https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8
//https://peter.bourgon.org/blog/2016/07/11/context.html
//code:
//exploring context
//background
//https://play.golang.org/p/cByXyrxXUf
//WithCancel
//throwing away CancelFunc
//https://play.golang.org/p/XOknf0aSpx
//using CancelFunc
//https://play.golang.org/p/UzQxxhn_fm
//Example
//https://play.golang.org/p/Lmbyn7bO7e
//func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
//https://play.golang.org/p/wvGmvMzIMW
//cancelling goroutines with deadline
//func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
//https://play.golang.org/p/Q6mVdQqYTt
//with timeout
//func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
//https://play.golang.org/p/OuES9sP_yX
//with value
//func WithValue(parent Context, key, val interface{}) Context
//https://play.golang.org/p/8JDCGk1K4P
//video:  163

