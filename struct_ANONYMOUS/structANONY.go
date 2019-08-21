package main

import "fmt"

//type person struct {   //has a name
//	first string
//	last string
//	age int
//}

func main()  {
	p1 := struct {   //anonymous struct and composite literal
		first string  //anonymous because it has no name
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,

	}


	//p1 := person{  //composite literal
	//		first: "James",
	//		last: "Bond",
	//		age: 32,
	//}
	fmt.Println(p1)
}
