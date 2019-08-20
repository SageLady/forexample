package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	} //END map
	//Adding an element to map
	m[`ruth_dr`] = []string{`sex talk`, `old lady`, `funny`}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		} //END for
	} //END for
} //END main

//Hands-on exercise #9
//Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
//solution: https://play.golang.org/p/_CkxAhRrDJ
//video: 079
//Solution:
//package main

//import (
//"fmt"
//)

//func main() {
//	m := map[string][]string{
//		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
//		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
//		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
//	}

	// fmt.Println(m)

//	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

//	for k, v := range m {
//		fmt.Println("This is the record for", k)
//		for i, v2 := range v {
//			fmt.Println("\t", i, v2)
//		}
//	}
//}
