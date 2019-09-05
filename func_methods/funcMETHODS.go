package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}
//func (r receiver) identifier (parameters) (return(s)) { code }
//The receiver allows for the type to access to have access
//to the function. Hence the struc secretAgent has access to the
//function speak and remains in scope.


func (s secretAgent) speak()  {
		fmt.Printf("\nInside of speak with Secret Agent\t%v\t%v", s.first, s.last)
}

func main()  {
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

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
} //END main
