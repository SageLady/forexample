package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

//Keyword = var, identifier = x, type = int
//var x int

//Keyword = type, identifier = human, type = interface
type human interface {
	speak()
}

//A value can be of more than one type. This is because human can be of more than one type.
//Show below with sa1
func bar(h human)  {
	switch h.(type) {
	case person:
		fmt.Println("Type person", h.(person).first, h.(person).last)   //Example of assertion,
	case secretAgent:
		fmt.Println("Type secretAgent", h.(secretAgent).first, h.(secretAgent).last)
	default:
		fmt.Println("No type")
	}
		fmt.Println("\nI was passed into bar", h)
}

func (s secretAgent) speak()  {
	fmt.Printf("\nInside of speak with Secret Agent %v\t%v - the secret agent speak\n", s.first, s.last)
}

func (p person) speak()  {
	fmt.Printf("\nInside of speak with Secret Agent %v\t%v - the person speak\n", p.first, p.last)
}

type hotdog int
func main()  {
	// Value sa1 is of type secretAgent, but because it also has speak attached to it it is also
	//type human
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk:    true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk:    true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}


	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)
	// Conversion:
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

} //END main
