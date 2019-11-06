//package main

//import (
//	"encoding/json"
//	"fmt"
//	"log"
//)

//type person struct {
//	First   string
//	Last    string
//	Sayings []string
//}

//func main() {
//	p1 := person{
//		First:   "James",
//		Last:    "Bond",
//		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
//	}

//	bs, _ := json.Marshal(p1)
//	fmt.Println(string(bs))

//}

//OUTPUT:
//C:\Users\rsouza\AppData\Local\Temp\___go_build_main_go.exe #gosetup
//{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes
//?","Never say never"]}

//Process finished with exit code 0



//Hands-on exercise #1
//Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.
//solution:
//https://play.golang.org/p/tn8oiuL1Yn
//video: 176

//SOLUTION:
package main

import (
"encoding/json"
"fmt"
"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}