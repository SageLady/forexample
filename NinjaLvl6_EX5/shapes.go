package main

import (
	"fmt"
	"math"
)


type SQUARE struct {
	width float64
	length float64
}

type CIRCLE struct {
	radius float64
}

func (cArea CIRCLE) Area() float64{
	var AREA float64
	AREA = math.Pi * (cArea.radius * 2)
	return AREA
}

func (sArea SQUARE) Area() float64{
	var AREA float64
	//sArea.width = 2.000
	AREA = sArea.length * sArea.width
	return AREA
}

type shape interface {
	Area() float64
}

func info(s shape) {
	fmt.Println(s.Area())
}

func main()  {
s1 := SQUARE{
	width:  3.40,
	length: 2.56,
}

c1 := CIRCLE{
	radius: 4.56,
		}

info(c1)
info(s1)


} // END main

//Hands-on exercise #5
//create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//circle area= π r 2
//square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle
//code: https://play.golang.org/p/NGGikHNvHv
//video: 106

//SOLUTION
//package main

//import (
//"fmt"
//"math"
//)

//type circle struct {
//	radius float64
//}

//type square struct {
//	length float64
//}

//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}

//func (s square) area() float64 {
//	return s.length * s.length
//}

//type shape interface {
//	area() float64
//}

//func info(s shape) {
//	fmt.Println(s.area())
//}

//func main() {
//	circ := circle{
//		radius: 12.345,
//	}

//	squa := square{
//		length: 15,
//	}

//	info(circ)
//	info(squa)
//}

