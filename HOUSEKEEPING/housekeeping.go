package main

import "fmt"

//Bill Kennedy teaches online intermediate classes of golang
//We create VALUES of a certain type that are stored in VARIABLES
//and those VARIABLES have identifiers
var x int //static type int, var x is of type int, x is identifier
type person struct{  //var type are keywords plus identifiers, person is identifier
	first string
	last string
}

type foo int
var y foo     //alias type

const bar int  = 42    //____ indentifier and type
func main()  {
	p1 := person{ //composite literal
		first: "James",
		last: "bond",
	}
	fmt.Println(p1)
	y = 42   //indeterminate via foo compiler need to determine
	        //base type. Contstant of a kind
	fmt.Printf("%T\n", int(y))  //using conversion on y identifier
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)
}
//Is go an Object-oriented language?
//Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).

//Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.
