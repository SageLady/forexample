package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age   int
}

func (p person)String() string {
		return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

//WITHOUT String()
//[{James 32} {Moneypenny 27} {Q 64} {M 56}]
//[{Moneypenny 27} {James 32} {M 56} {Q 64}]
//[{James 32} {M 56} {Moneypenny 27} {Q 64}]

//WITH String()
//[James: 32 Moneypenny: 27 Q: 64 M: 56]
//[Moneypenny: 27 James: 32 M: 56 Q: 64]
//[James: 32 M: 56 Moneypenny: 27 Q: 64]



type ByAge []person
func (a ByAge) Len() int{
		return len(a) }
func (a ByAge) Swap(i, j int) {
		a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool{
 		return a[i].Age < a[j].Age }

type ByName []person

func (n ByName) Len() int{
	return len(n) }
func (n ByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool{
	return n[i].Name < n[j].Name }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}   //people is a slice of person

	fmt.Println(people)   //OUTPUT: [James: 32 Moneypenny: 27 Q: 64 M: 56]

	//ByAge implements sort.
	// Interface for []Person based on the
	//age field
	sort.Sort(ByAge(people))
	fmt.Println(people)   //OUTPUT: //[Moneypenny: 27 James: 32 M: 56 Q: 64]
	//ByName implements sort.
	// Interface for []Person based on the
	//name field
	sort.Sort(ByName(people))
	fmt.Println(people)    //OUTPUT: [James: 32 M: 56 Moneypenny: 27 Q: 64]
}

//OUTPUT:  (ByAge added)
//[James: 32 Moneypenny: 27 Q: 64 M: 56]
//[Moneypenny: 27 James: 32 M: 56 Q: 64]

//OUTPUT: (ByName added)
//[James: 32 Moneypenny: 27 Q: 64 M: 56]
//[Moneypenny: 27 James: 32 M: 56 Q: 64]
//[James: 32 M: 56 Moneypenny: 27 Q: 64]


//Original code:
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

//OUTPUT to ORIGINAL CODE:
//[{James 32} {Moneypenny 27} {Q 64} {M 56}]


