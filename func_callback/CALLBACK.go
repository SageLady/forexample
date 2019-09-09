package main

import (
"fmt"
)

func main()  {
	s := sum(1,2,3,4,5,6,7)
	fmt.Println(s)
	ii := []int{1,2,3,4,5,6,7,8,9,10, 11,12,13,14,14,15,16}
	t := sum(ii...)
	fmt.Println("All numbers",t)
	s2 := even(sum,ii...)
	fmt.Println("even numbers", s2)
	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)

}  //END main

func sum(xi...int) int { // variadic parameter
	total := 0
	for _, v := range  xi{
		total += v   // += increments variable the amount. So, if total equals 2 and value is 2, the += would make total equal 4.
	}
	return total
}

func even(f func(xi...int) int, vi...int) int {
var yi [] int
for _, v := range vi {
if v % 2 == 0 {
yi = append(yi, v)
} //END if
} //END for
return f(yi...)
} //END even

func odd(f func(xi...int) int, vi...int) int {
	var yi [] int
	for _, v := range vi {
		if v % 2 != 0 {
			yi = append(yi, v)
		} //END if
	} //END for
	return f(yi...)
} //END odd