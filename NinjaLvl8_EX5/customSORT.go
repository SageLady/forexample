package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) String() string {
			return fmt.Sprintf("%s %s: %d", u.First, u.Last, u.Age)
}

type ByAge []user
func (a ByAge) Len() int{
	return len(a)
}    // END Len
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}	// END Swap
func (a ByAge) Less(i, j int) bool{
	return a[i].Age < a[j].Age
}  // END Less

type ByLast []user

func (l ByLast ) Len() int{
		return len(l)
}  // END Len
func (l ByLast) Swap(i, j int) {
		l[i], l[j] = l[j], l[i]
}	// END Swap
func (l ByLast) Less(i, j int) bool{
		return l[i].Last < l[j].Last
}  // END Less

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)
	// your code goes here
	fmt.Println("         ----UNSORTED----")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
			for _, v := range u.Sayings {
				fmt.Println("\t", v)
			}  // END for
		}  // END for
     fmt.Println("         ----SORT BY AGE----")
	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}  // END for
	}  // END for
	//fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println("         ----SORT BY LAST NAME----")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}  // END for
	}  // END for
}

//Hands-on exercise #5
//Starting with this code, sort the []user by
//age
//last
//Also sort each []string “Sayings” for each user
//print everything in a way that is pleasant
//solution: https://play.golang.org/p/8RKkdtLl6w

//STARTING CODE:
//package main

//import (
//"fmt"
//)

//type user struct {
//	First   string
//	Last    string
//	Age     int
//	Sayings []string
//}

//func main() {
//	u1 := user{
//		First: "James",
//		Last:  "Bond",
//		Age:   32,
//		Sayings: []string{
//			"Shaken, not stirred",
//			"Youth is no guarantee of innovation",
//			"In his majesty's royal service",
//		},
//	}

//	u2 := user{
//		First: "Miss",
//		Last:  "Moneypenny",
//		Age:   27,
//		Sayings: []string{
//			"James, it is soo good to see you",
//			"Would you like me to take care of that for you, James?",
//			"I would really prefer to be a secret agent myself.",
//		},
//	}

//	u3 := user{
//		First: "M",
//		Last:  "Hmmmm",
//		Age:   54,
//		Sayings: []string{
//			"Oh, James. You didn't.",
//			"Dear God, what has James done now?",
//			"Can someone please tell me where James Bond is?",
//		},
//	}

//	users := []user{u1, u2, u3}
//
//	fmt.Println(users)

	// your code goes here

//}

//SOLUTION:
//package main

//import (
//"fmt"
//"sort"
//)

//type user struct {
//	First   string
//	Last    string
//	Age     int
//	Sayings []string
//}

//type ByAge []user

//func (a ByAge) Len() int           { return len(a) }
//func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//type ByLast []user

//func (l ByLast) Len() int           { return len(l) }
//func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
//func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

//func main() {
//	u1 := user{
//		First: "James",
//		Last:  "Bond",
//		Age:   32,
//		Sayings: []string{
//			"Shaken, not stirred",
//			"Youth is no guarantee of innovation",
//			"In his majesty's royal service",
//		},
//	}

//	u2 := user{
//		First: "Miss",
//		Last:  "Moneypenny",
//		Age:   27,
//		Sayings: []string{
//			"James, it is soo good to see you",
//			"Would you like me to take care of that for you, James?",
//			"I would really prefer to be a secret agent myself.",
//		},
//	}

//	u3 := user{
//		First: "M",
//		Last:  "Hmmmm",
//		Age:   54,
//		Sayings: []string{
//			"Oh, James. You didn't.",
//			"Dear God, what has James done now?",
//			"Can someone please tell me where James Bond is?",
//		},
//	}
//
//	users := []user{u1, u2, u3}
//	for _, u := range users {
//		fmt.Println(u.First, u.Last, u.Age)
//		sort.Strings(u.Sayings)
//		for _, v := range u.Sayings {
//			fmt.Println("\t", v)
//		}
//	}

//	fmt.Println("-------")
//
//	sort.Sort(ByAge(users))
//	for _, u := range users {
//		fmt.Println(u.First, u.Last, u.Age)
//		for _, v := range u.Sayings {
//			fmt.Println("\t", v)
//		}
//	}

//	fmt.Println("-------")

//	sort.Sort(ByLast(users))
//	for _, u := range users {
//		fmt.Println(u.First, u.Last, u.Age)
//		for _, v := range u.Sayings {
//			fmt.Println("\t", v)
//		}
//	}
//
//}
