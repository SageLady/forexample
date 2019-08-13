package main
//Code prints out only even numbers
//up to 100
import "fmt"

func main()  {
	x := 1
	for  {
		x++
		if  x > 100 {
				break
		}

		if x%2 != 0 {
				continue
		}
		fmt.Println(x)
	}
}
