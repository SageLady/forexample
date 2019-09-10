package main

import "fmt"

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9,10, 11,12,13,14,14,15,16}
	t := foo(ii...)
	fmt.Println("foo-Sum of all numbers",t)

	vi := []int{45,45,50,100,100}
	u := bar(vi)
	fmt.Println("bar-Sum of all numbers",u)

}

func foo(xi...int) int{
	total := 0
	for _,v := range xi {
		total += v
	}
	return total
}

func bar(vi []int) int{
	sum := 0
	for _, v := range vi {
		sum += v
	}
	return sum
}
//Hands-on exercise #2
//create a func with the identifier foo that
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in
//create a func with the identifier bar that
//takes in a parameter of type []int
//returns the sum of all values of type int passed in
//code: https://play.golang.org/p/B0yRxtBQPD
//video: 103



