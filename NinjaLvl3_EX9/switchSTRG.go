package main

import "fmt"

func main()  {
	favSport := "SURFING"
	switch favSport {
	case "TENNIS":
		fmt.Println("TENNIS")
	case "GOLF":
		fmt.Println("GOLF")
	case "SOCCER":
		fmt.Println("SOCCER")
	case "SURFING":
		fmt.Println("got to go to Hawaii")
	}

}


//Hands-on exercise #9
//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
//solution: https://play.golang.org/p/h-8FnjpzWB
//video: 058

//SOLUTION:
//package main

//import (
//"fmt"
//)

//func main() {
//	favSport := "surfing"
//	switch favSport {
//	case "skiing":
//		fmt.Println("go to the mountains!")
//	case "swimming":
//		fmt.Println("go to the pool!")
//	case "surfing":
//		fmt.Println("go to hawaii!")
//	case "wingsuit flying":
//		fmt.Println("what would you like me to say at your funeral?")
//	}
//}
