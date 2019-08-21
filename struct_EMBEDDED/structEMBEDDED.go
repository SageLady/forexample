package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	first string
	ltk bool
}

func main()  {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 22,
		},
		first: "ALIAS NAME",
		ltk:    true,
	}

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
	fmt.Println(sa1)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1.ltk, sa1.age, sa1.first, sa1.person.first, sa1.last)
	fmt.Println(p1.age, p1.last)
	fmt.Println(p2.age, p2.last)
}