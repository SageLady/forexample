package main

import "fmt"

type person struct {
	first string
	last string
	favorite []string
}

func main()  {
	p1 := person{
		first: "Charles",
		last:  "Williams",
		favorite: []string {
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favorite {
			fmt.Println(i, val)
		}
		fmt.Println("-------------")
	} //END for
} //END main

//Hands-on exercise #2
//Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
//solution: https://play.golang.org/p/RvvLyAMvGo
//video: 087
//SOLUTIONS:
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

//	m := map[string]person{
///		p1.last: p1,
//		p2.last: p2,
//	}

//	for _, v := range m {
//		fmt.Println(v.first)
//		fmt.Println(v.last)
//		for i, val := range v.favFlavors {
//			fmt.Println(i, val)
//		}
//		fmt.Println("-------")
//	}
//}



