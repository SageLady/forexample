package main
//Finding sort methods:
//1. Go to golang.org
//2. Select "Documents" link
//3. Go to "References-->Package Documentation"
//4. Ctrl+f to search "sort"
//5. Boom you are taken to 'sort' hyperlink and definition
//6. Instructor prefers: https://godoc.org/sort

import (
	"fmt"
	"sort"
)

func main()  {
	xi := []int{4, 7, 3, 52, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)
	sort.Ints(xi)
	fmt.Println("Now xi is sorted:",xi)
	sort.Strings(xs)
	fmt.Println("Now the strings, xs, are sorted:", xs)


}

//OUTPUT:
//[4 7 3 52 99 18 16 56 12]
//[James Q M Moneypenny Dr. No]
//Now xi is sorted: [3 4 7 12 16 18 52 56 99]
//Now the strings, xs, are sorted: [Dr. No James M Moneypenny Q]


//Class notes:
//Sort custom
//Here is how we sort on fields in a struct.
//code:
//starting code:
//https://play.golang.org/p/UhXN-G2FwY
//sorted
//by age: https://play.golang.org/p/kqmJovOU5V
//by name: https://play.golang.org/p/he70VcFmdM
//video: 123

//Starting Code:
//func main()  {
//xi := []int{4, 7, 3, 52, 99, 18, 16, 56, 12}
//xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

//fmt.Println(xi)
//fmt.Println(xs)
//}

//OUTPUT:
//[4 7 3 52 99 18 16 56 12]
//[James Q M Moneypenny Dr. No]



//package main

//import (
//"fmt"
//)

//type person struct {
//	first string
//	age   int
//}

//func main() {
//	p1 := person{"James", 32}
//	p2 := person{"Moneypenny", 27}
//	p3 := person{"Q", 64}
//	p4 := person{"M", 56}

//	people := []person{p1, p2, p3, p4}

//	fmt.Println(people)
//}