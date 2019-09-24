package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}




//Hands-on exercise #4
//Using STARTING CODE, below, sort the []int and []string for each person.
//solution: https://play.golang.org/p/jz_llY1aPp
//video: 128

//STARTING CODE:
//package main

//import (
//"fmt"
//)

//func main() {
//	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
//	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

//	fmt.Println(xi)
	//Note: sort xi
//	fmt.Println(xi)

//	fmt.Println(xs)
	//Note: sort xs
//	fmt.Println(xs)
//}

//SOLUTION:
//package main

//import (
//"fmt"
//"sort"
//)

//func main() {
//	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
//	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

//	fmt.Println(xi)
//	sort.Ints(xi)
//	fmt.Println(xi)

//	fmt.Println(xs)
//	sort.Strings(xs)
//	fmt.Println(xs)
//}

