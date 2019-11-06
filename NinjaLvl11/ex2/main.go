package main

import (
	"encoding/json"
	"fmt"
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

	bs, err := toJSON(p1)

	fmt.Println(string(bs))
	fmt.Println(err)

}

         // toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
	}
	return bs, nil
}

//        // toJSON needs to return an error also
//func toJSON(a interface{}) ([]byte, error) {
//	bs, err := json.Marshal(a)
//	if err != nil {
//		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
//	}
//	return bs, nil
//}


//Hands-on exercise #2
//Start with this code. Create a custom error message using “fmt.Errorf”.
//solution:
//https://play.golang.org/p/HugU4HJEEO  - ONE
//https://play.golang.org/p/NII-lmGasj  - TWO
//https://play.golang.org/p/Vo5kIoR-sG  - THREE
//video: 177

//SOLUTION ONE:
//package main

//import (
//"encoding/json"
//"fmt"
//"log"
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

//	bs, err := toJSON(p1)
//	if err != nil {
//		log.Fatalln(err)
//	}

//	fmt.Println(string(bs))
//}

//        // toJSON needs to return an error also
//func toJSON(a interface{}) ([]byte, error) {
//	bs, err := json.Marshal(a)
//	if err != nil {
//		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
//	}
//	return bs, nil
//}

//SOLUTION TWO:
//package main

//import (
//"encoding/json"
//"fmt"
//"log"
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

//	bs, err := toJSON(p1)
//	if err != nil {
//		log.Fatalln(err)
//	}

//	fmt.Println(string(bs))

//}

//func toJSON(a interface{}) ([]byte, error) {
//	bs, err := json.Marshal(a)
//	if err != nil {
//		return []byte{}, fmt.Errorf("There was an error marshalling %v in toJSON: %v", a, err)
//	}
//	return bs, nil
//}

//SOLUTION THREE:
//package main

//import (
//"encoding/json"
//"errors"
//"fmt"
//"log"
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

//	bs, err := toJSON(p1)
//	if err != nil {
//		log.Fatalln(err)
//	}

//	fmt.Println(string(bs))
//}

// toJSON needs to return an error also
//func toJSON(a interface{}) ([]byte, error) {
//	bs, err := json.Marshal(a)
//	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
//		return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON %v", err))
//	}
//	return bs, nil
//}



