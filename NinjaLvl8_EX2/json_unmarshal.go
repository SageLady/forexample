package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Last string `json:"last"`
	Age int `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	bs := []byte(s)
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []user

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}  // END if
	fmt.Println("All of the people:", people)

	for i, v := range people {
			fmt.Println("\nPERSON NUMBER", i)
			fmt. Println(v.First, v.Last, v.Age)
			fmt. Println(v.Sayings)
	}  // END for


}  // END main




//Hands-on exercise #2
//Starting with "STARTING CODE" (below), unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/
//code:
//code setup (just fyi, not needed for exercise):
//https://play.golang.org/p/nWPP5Z2Q4e
//solution:
//https://play.golang.org/p/r8oSG8DIPR
//video: 126

//STARTING CODE:
//package main

//import (
//	"fmt"
//)

//func main() {
//	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
//	fmt.Println(s)

//}

//SOLUTION:
//package main

//import (
//"encoding/json"
//"fmt"
//)

//type person struct {
//	First   string   `json:"First"`
//	Last    string   `json:"Last"`
//	Age     int      `json:"Age"`
//	Sayings []string `json:"Sayings"`
//}

//func main() {
//	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
//	fmt.Println(s)

//	var people []person

//	err := json.Unmarshal([]byte(s), &people)
//	if err != nil {
//		fmt.Println(err)
//	}

//	fmt.Println(people)

//	for i, person := range people {
//		fmt.Println("Person #", i)
///		fmt.Println("\t", person.First, person.Last, person.Age)
	//	for _, saying := range person.Sayings {
//			fmt.Println("\t\t", saying)
//		}
//	}
//}

//OUTPUT OF SOLUTION:
//[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]
//[{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]
//Person # 0
//James Bond 32
//Shaken, not stirred
//Youth is no guarantee of innovation
//In his majesty's royal service
//Person # 1
//Miss Moneypenny 27
//James, it is soo good to see you
//Would you like me to take care of that for you, James?
//I would really prefer to be a secret agent myself.
//Person # 2
//M Hmmmm 54
//Oh, James. You didn't.
//Dear God, what has James done now?
//Can someone please tell me where James Bond is?

