package main

import (
	"fmt"
)

type person struct {
	name string
}
//SOLUTION:
//type person struct {
//	first string
//}

func (p *person) speak()  {
	fmt.Println("Hello World!")

}
//SOLUTION:
//func (p *person) speak() {
//	fmt.Println("Hello")
//}

type human interface {    //interface
	speak()
}
//SOLUTION:
//type human interface {
//	speak()
//}

func saySomething(h human)  {    //implicitly implement interface
	h.speak()
}
//SOLUTION:
//func saySomething(h human) {
//	h.speak()
//}

func main()  {
	p1 := person{"Frank Charlie",}
	fmt.Printf("Person: %v says:", p1.name)

	saySomething(&p1)
	p1.speak()
}
//func main() {
//	p1 := person{
//		first: "James",
//	}

//	saySomething(&p1)

	// does not work
	// saySomething(p1)

//	p1.speak()
//}


//Hands-on exercise #2
//This exercise will reinforce our understanding of method sets:
//  3. create a type person struct
//	9. attach a method speak to type person using a pointer receiver
//	9. *person
//	13. create a type human interface
//	13. to implicitly implement the interface, a human must have the speak method
//	17. create func “saySomething”
//	17. have it take in a human as a parameter
//	17. have it call the speak method
//	show the following in your code
//	you CAN pass a value of type *person into saySomething
//	you CANNOT pass a value of type person into saySomething
//	here is a hint if you need some help
//	https://play.golang.org/p/FAwcQbNtMG

//	Receivers       Values
//	-----------------------------------------------
//	(t T)           T and *T
//	(t *T)          *T

//	code: https://github.com/GoesToEleven/go-programming
//	video: 149


//HINT:
//package main

//import (
//"fmt"
//"math"
//)

//type circle struct {
//	radius float64
//}

//type shape interface {
//	area() float64
//}

//func (c *circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}

//func info(s shape) {
//	fmt.Println("area", s.area())
//}

//func main() {
//	c := circle{5}
	//Note: info(c)
	//info(&c)
	//Note: fmt.Println(c.area())
//}

//SOLUTION:
//package main

//import "fmt"

//type person struct {
//	first string
//}

//type human interface {
//	speak()
//}

//func (p *person) speak() {
//	fmt.Println("Hello")
//}

//func saySomething(h human) {
//	h.speak()
//}

//func main() {
//	p1 := person{
//		first: "James",
//	}

//	saySomething(&p1)

	// does not work
	// saySomething(p1)   //commented out: doesn't work

//	p1.speak()   //works
//}