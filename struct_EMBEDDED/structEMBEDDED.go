package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main()  {
	p1 := person {
		first: "James",
		last: "Bond",
		age: 22,
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 21,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.age, p1.last)
	fmt.Println(p2.age, p2.last)
}