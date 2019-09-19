package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int
}

func main()  {
	p1 := person {
		First: "James",
		Last: "Bond",
		Age: 32,
	}

	p2 := person {
		First: "Miss",
		Last: "Moneypenny",
		Age: 32,
	}
	people := []person{p1, p2,}
	fmt.Println(people)

//Need to send person the data, so
//use json
bs, err := json.Marshal(people)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(bs))

}

//Example of struct fields being lowercase
//type person struct {
//	first string
//	last string
//	age int
//}
//OUTPUT #1:
//[{James Bond 32} {Miss Moneypenny 32}]  //GO FORMAT
//[{},{}]  // Marshalling looked at the case of the struct vars

//After changing all the struct field names to capital letters
//then output 2/
//OUTPUT #2:
//[{James Bond 32} {Miss Moneypenny 32}]   //GO FORMAT
//[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","A
//ge":32}]    //MARSHAL FORMAT (a string)

