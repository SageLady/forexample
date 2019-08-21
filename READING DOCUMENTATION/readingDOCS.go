package main

import "fmt"

type person struct {
	first string
	last string
	age int
}
type secretAgent struct {
	person   //aka anonymous field  or embedded field OR unqualified
	         //field name acts as the field name
	first string
	ltk bool
}

var x, y int

func main()  {
	x = 42
	y = 43
	fmt.Println(x,y)

	x, y = y, x
	fmt.Println(x,y)

	sa1 := secretAgent{ //embedded field is person struct fields
		person: person{ //unqualified field name acting as field name
			first: "James",  //promoted field
			last: "Bond",  //promoted field
			age: 22,  //promoted field
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
	fmt.Println(p2)   //promoted fields are called directly from sa1
	fmt.Println(sa1.ltk, sa1.age, sa1.first, sa1.person.first, sa1.last)
	fmt.Println(p1.age, p1.last)
	fmt.Println(p2.age, p2.last)
}
