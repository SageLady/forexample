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
//Struct
//A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). Structs allow us to compose together values of different types.
//code: https://play.golang.org/p/hNI_rSK-C6
//video:081
