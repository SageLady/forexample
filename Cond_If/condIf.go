package main

import "fmt"

func main()  {
	x := 43

	if x == 40 {
		fmt.Println("Our value was 40")
    } else if x == 41 {
		fmt.Println("Our value was 41")
	} else if x == 42 {
		fmt.Println("Our value was 42")
	} else {
		fmt.Println("Our value was not 40, 41 or 42")
	}
}

//Initialization statement & scope of x in the if block only
//func main()  {
	//if x := 42; x == 2 {
	//	fmt.Println("001")
	//}
	//fmt.Println("here's a statement");fmt.Println("and another")
//}




	//if 2 == 2 {     //Expression statement
	//	fmt.Println("005")
	//}

	//if 2 != 2 {		//Expression statement
	//	fmt.Println("006")
	//}

	//if !(2 == 2) {		//Expression statement
	//	fmt.Println("007")
	//}

	//if !(2 != 2) {		//Expression statement
	//	fmt.Println("008")
	//}

	//if true {     //predeclared constant
	//	fmt.Println("001")
	//}

	//if false {		//predeclared constant
	//	fmt.Println("002")
	//}

	//if !true {		//predeclared constant
	//	fmt.Println("003")
	//}

	//if !false {		//predeclared constant
	//	fmt.Println("004")
	//}

//}

//Conditional - if statment
//If statement
//1. bool ( predeclared)
//1.a true
//1.b false
//2. the not operator !
//3. initialization statement
//4. if / else
//5. if / else if / else
//6. if / else if / else if / ... / else