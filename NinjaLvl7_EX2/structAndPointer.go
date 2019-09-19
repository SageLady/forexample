package main

import "fmt"

type person struct {
	first string
	last string
	address string
}

func changeMe(p *person,fname string, lname string, address string)  {
//	p.first = fname
	(*p).first = fname
	p.last = lname
	p.address = address
	fmt.Println("Updates inside of changeMe:", p.first, p.last, p.address)
}

func main()  {
	p1 := person{
		first:   "First",
		last:    "Last",
		address: "2 A Street City State Zip",
	}
	fmt.Println(p1)
	changeMe(&p1, "Harold", "Kumar", "2 B Street, Hayward CA 94545")
	fmt.Println("Updates outside of changeMe:", p1.first, p1.last, p1.address)
	
}


//Hands-on exercise #2
//create a person struct
//create a func called “changeMe” with a *person as a parameter
//change the value stored at the *person address
//important
//to dereference a struct, use (*value).field
//p1.first
//(*p1).first
//“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
//https://golang.org/ref/spec#Selectors
//create a value of type person
//print out the value
//in func main
//call “changeMe”
//in func main
//print out the value
//code: https://play.golang.org/p/AehM662HKS
//video: 117

//SOLUTION:
//package main

//import (
//"fmt"
//)

//type person struct {
//	name string
//}

//func main() {
//	p1 := person{
//		name: "James Bond",
//	}
//	fmt.Println(p1)
//	changeMe(&p1)
//	fmt.Println(p1)
//}

//func changeMe(p *person) {
//	p.name = "Miss Moneypenny"
	// (*p).name = "Miss Moneyp"
//}
