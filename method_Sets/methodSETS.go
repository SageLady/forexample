package main

func main()  {
	
}
//Method sets
//Method sets determine what methods attach to a TYPE. It is exactly what the name says: What is the set of methods for a given type? That is its method set.

//IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....”
//~ golang spec
//The above “important” was left out of this video but will be discussed in the “concurrency” section in a video called “method sets revisited”.

//a NON-POINTER RECEIVER
//works with values that are POINTERS or NON-POINTERS.
//a POINTER RECEIVER
//only works with values that are POINTERS.

//Receivers       Values
//-----------------------------------------------
//(t T)           T and *T
//(t *T)          *T

//code:
//NON-POINTER RECEIVER & NON-POINTER VALUE
//https://play.golang.org/p/2ZU0QX12a8
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

//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}

//func info(s shape) {
//	fmt.Println("area", s.area())
//}

//func main() {
//	c := circle{5}
//	info(c)
//}

//NON-POINTER RECEIVER & POINTER VALUE
//https://play.golang.org/p/glWZmm0gY6
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

//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}

//func info(s shape) {
//	fmt.Println("area", s.area())
//}

//func main() {
//	c := circle{5}
//	info(&c)
//}

//POINTER RECEIVER & POINTER VALUE
//https://play.golang.org/p/pWFxsg6MSe
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
//	info(&c)
//}

// EXAMPLE 1: POINTER RECEIVER & NON-POINTER VALUE
// *NOTE: CODE DOES NOT RUN
//https://play.golang.org/p/G3lEy-4Mc8
//ERROR OUTPUT:
//./prog.go:26:6: cannot use c (type circle) as type shape in argument to info:
//circle does not implement shape (area method has pointer receiver)

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
//	info(c)
//}

// EXAMPLE 2: POINTER RECEIVER & NON-POINTER VALUE
// This codes does run - notice the difference -  method set determines the INTERFACES that the type implements
//https://play.golang.org/p/KK8gjsAWBZ
//video: 115

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
	//Originally commented out by teacher:  info(c)
//	fmt.Println(c.area())
//}

