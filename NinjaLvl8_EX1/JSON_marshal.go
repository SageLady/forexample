package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}





//Hands-on exercise #1
//Starting with "STARTING CODE" (shown below), marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
//solution: https://play.golang.org/p/8BK6PXj3aG
//video: 125

//STARTING CODE:
//package main

//import (
//"fmt"
//)

//type user struct {
//	First string
//	Age   int
//}

//func main() {
//	u1 := user{
//		First: "James",
//		Age:   32,
//	}

//	u2 := user{
//		First: "Moneypenny",
//		Age:   27,
//	}

//	u3 := user{
//		First: "M",
//		Age:   54,
//	}

//	users := []user{u1, u2, u3}

//	fmt.Println(users)

	// your code goes here
//}


//SOLUTION
//package main

//import (
//"encoding/json"
//"fmt"
//)

//type user struct {
//	First string
//	Age   int
//}

//func main() {
//	u1 := user{
//		First: "James",
//		Age:   32,
//	}

//	u2 := user{
//		First: "Moneypenny",
//		Age:   27,
//	}

//	u3 := user{
//		First: "M",
//		Age:   54,
//	}

//	users := []user{u1, u2, u3}

//	fmt.Println(users)

//	bs, err := json.Marshal(users)
//	if err != nil {
//		fmt.Println(err)
//	}

//	fmt.Println(string(bs))
//}
