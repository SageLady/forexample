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

}

// toJSON needs to return an error also
func toJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
}


//Hands-on exercise #2
//Start with this code. Create a custom error message using “fmt.Errorf”.
//solution:
//https://play.golang.org/p/HugU4HJEEO
//https://play.golang.org/p/NII-lmGasj
//https://play.golang.org/p/Vo5kIoR-sG
//video: 177
