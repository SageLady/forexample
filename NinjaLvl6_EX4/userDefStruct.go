package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak()  {
	fmt.Printf("My name is %v %v", p.first, p.last)
	fmt.Printf("\nI am %v years old", p.age)
}


func main()  {
p1 := person{
	first: "Freddee",
	last:  "Messer",
	age:   1,
}

p1.speak()
}

//Hands-on exercise #4
//Create a user defined struct with
//the identifier “person”
//the fields:
//first
//last
//age
//attach a method to type person with
//the identifier “speak”
//the method should have the person say their name and age
//create a value of type person
//call the method from the value of type person
//code: https://play.golang.org/p/NnXyWdqbbw
//video: 105

//SOLUTION
//package main

//import (
//"fmt"
//)

//type person struct {
//	first string
//	last  string
//	age   int
//}

//func (p person) speak() {
//	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
//}

//func main() {
//	p1 := person{
//		first: "James",
//		last:  "Bond",
//		age:   32,
//	}

//	p1.speak()
//}
