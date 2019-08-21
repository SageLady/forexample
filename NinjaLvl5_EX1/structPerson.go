package main

import "fmt"

type person struct {
	first string
	last string
	favorite []string
}
func main()  {
	p1 := person{
		first:    "bob",
		last:     "williams",
		favorite: []string{
				"chocolate",
				"bubble gum",
				"pistachio",
		},
	}
	p2 := person{
		first:    "michelle",
		last:     "franks",
		favorite: []string{
				"peach",
				"sherbert",
				"pina colata",
		},
	}
	fmt.Printf("The record for %v %v\n", p1.first, p1.last)
	for i, v := range p1.favorite {
			fmt.Println(i,v)
	} //END for

	fmt.Printf("The record for %v %v\n", p2.first, p2.last)
	for i, v := range p2.favorite {
		fmt.Println(i,v)
	} //END for
} //END main

//Exercises - Ninja Level 5
//Hands-on exercise #1
//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
//solution:
//https://play.golang.org/p/ouRHmH_POg
//video: 086
//SOLUTION:
//package main

//import (
//"fmt"
//)

//type person struct {
//	first      string
//	last       string
//	favFlavors []string
//}

//func main() {
//	p1 := person{
//		first: "James",
//		last:  "Bond",
//		favFlavors: []string{
//			"chocolate",
//			"martini",
//			"rum and coke",
//		},
//	}

//	p2 := person{
//		first: "Miss",
//		last:  "Moneypenny",
//		favFlavors: []string{
//			"strawberry",
//			"vanilla",
//			"capuccino",
//		},
//	}

//	fmt.Println(p1.first)
//	fmt.Println(p1.last)
//	for i, v := range p1.favFlavors {
//		fmt.Println(i, v)
//	}

//	fmt.Println(p2.first)
//	fmt.Println(p2.last)
//	for i, v := range p2.favFlavors {
//		fmt.Println(i, v)
//	}
//}

