package main

//RESOURCE: the teacher's code for web development
//https://github.com/GoesToEleven/golang-web-dev/tree/master/040_json
//https://github.com/GoesToEleven/golang-web-dev/blob/master/040_json/README.html
//rawgit = end of life because malware and misuse of website. Closed down 10-2019.
//https://rawgit.com/
//https://github.com/GoesToEleven/golang-web-dev
//CONVERTING JSON to GO
//https://mholt.github.io/json-to-go/
//Marshal and unmarshal used a lot in web development

import (
	"encoding/json"
	"fmt"
)
//MARSHALLING CODE takes GO and turns it into JSON
// START ORIGINAL STRUCT CODE: marshalling, struct
//type person struct {
//	First string
//	Last string
//	Age int
//}
// END ORIGINAL STRUCT CODE: marshalling, struct

type person struct {
	First string `json:"first"`
	Last string `json:"last"`
	Age int `json:"age"`
}

func main() {

	//UNMARSHALLING CODE takes JSON and converts to GO
	//s = JSON  (Note: no \n or carriage returns can be in string
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":32}]`
	bs := []byte(s) //bs = byte slice

	fmt.Printf("%T\n", s)   //OUTPUT: string
	fmt.Printf("%T\n", bs)   //OUTPUT: uint8 = the set of all unsigned 8-bit integers (0 to 255)
                             //byte is an alias for uint8
    //people := []person{}    //This code equals the code on next line.
    var people []person       //Preferred by Todd McCloud for readibility

	err := json.Unmarshal(bs,&people)
	if err != nil {
		fmt.Println(err)
	}   // END if
	fmt.Println("all of the data", people)   //OUTPUT: all of the data [{James Bond 32} {Miss Moneypenny 32}]

	for i, v := range people {
			fmt.Println("\nPERSON NUMBER", i)
			fmt.Println(v.First, v.Last, v.Age)
	}  //END for

	//for OUTPUT:
	//PERSON NUMBER 0
	//James Bond 32

	//PERSON NUMBER 1
	//Miss Moneypenny 32


} // END main





//START: ORIGINAL CODE: see marshalling
//		p1 := person {
//		First: "James",
//		Last: "Bond",
//		Age: 32,
//	}

//	p2 := person {
//		First: "Miss",
//		Last: "Moneypenny",
//		Age: 32,
//	}
//	people := []person{p1, p2,}
//	fmt.Println(people)

	//Need to send person the data, so
	//use json
//	bs, err := json.Marshal(people)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(bs))
//END ORIGINAL CODE: Marshalling

//https://documentation.progress.com/output/ua/OpenEdge_latest/index.html#page/dvjsn%2Fjson-data-types.html%23
//JSON primitive data types: string, number (1754.350.9682e-42) boolean, null
//JSON complex data types:
// Object - comma-delimited list of named values, either simple or complex
//enclosed in braces
//Array - a comma-delimited list of unamed values, either simple or complex, enclosed in braces

//Instructor's Notes:
//JSON unmarshal
//We can take JSON and bring it back into our Go program by unmarshalling that JSON. Great resources for understanding and working with JSON are shared.
//cool websites:
//http://rawgit.com/    NOTE: website closing down October 2019
//https://mholt.github.io/json-to-go/
//https://github.com/goestoeleven
//code:
//understanding JSON rawgit HTML page
//https://play.golang.org/p/O9q0DmpzsZ
//video: 120

