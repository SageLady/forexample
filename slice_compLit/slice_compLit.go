package main

import "fmt"

func main()  {
//	x := type{values}  //composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}

//a SLICE allows you to group together
//VALUES of the same TYPE

//COMPOSITE LITERALS:
//Composite literals construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated.
//The LiteralType's underlying type must be a struct, array, slice, or map type (the grammar enforces this constraint except when the type is given as a TypeName). The types of the elements and keys must be assignable to the respective field, element, and key types of the literal type; there is no additional conversion.
//