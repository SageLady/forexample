package main

import "fmt"

func main()  {
	//missing switch expression defaults to 'true'
	switch  {
	case false:
		fmt.Println("the false case")
	case true:
		fmt.Println("the true case")
	} //END - switch

	//* -------------------
	//switch on value:
	//And multiple cases per case
	//v := "Bond"
	//switch v {
	//case "Moneypenny", "Bond", "Dr No":
	//	fmt.Println("miss money or Bond or Dr No")
	//case "M":
	//	fmt.Println("This is M")
	//case "Q":
	//	fmt.Println("This is Q")
	//default:
	//	fmt.Println("The default case")
	//} //END - switch

	//switch on value:
	//Can either use literal or variable
	//n := "Bond"
	//switch n {
	//case "Moneypenny":
	//	fmt.Println("miss money")
	//case "Bond":
	//	fmt.Println("James Bond")
	//case "Q":
	//	fmt.Println("This is Q")
	//default:
	//	fmt.Println("The default case")
	//} //END - switch

	//switch using literal
	//switch "Bond" {
	//case "Moneypenny":
	//	fmt.Println("miss money")
	//case "Bond":
	//	fmt.Println("James Bond")
	//case "Q":
	//	fmt.Println("This is Q")
	//default:
	//	fmt.Println("The default case")
	//} //END - switch
	//* -------------------
	//default case basic:
	//switch  {
	//case false:
	//	fmt.Println("this should not print")
	//case (2 == 4):
	//	fmt.Println("this should not print2")
	//case (3 == 5):
	//	fmt.Println("this should not print3")
	//default:
	//	fmt.Println("The default case")
	//} //END - switch
	//* -------------------
	//Funky switch with fallthrough
	//it prints when case is false
	//switch  {
	//case false:
	//	fmt.Println("this should not print")
	//case (2 == 4):
	//	fmt.Println("this should not print2")
	//case (3 == 3):
	//	fmt.Println("print") //Without fallthrough case 3 prints only
	//	fallthrough
	//case (4 == 4):
	//	fmt.Println("also true, and does print")
	//	fallthrough
	//case (7 == 9):
	//	fmt.Println("not true 1")
	//	fallthrough
	//case (11 == 14):
	//	fmt.Println("not true 2")
	//	fallthrough
	//case (15 == 15):
	//	fmt.Println("not true 15")
	//} //END - switch
	//* -------------------
	//fallthrough example on the first true case
	//switch  {
	//case false:
	//	fmt.Println("this should not print")
	//case (2 == 4):
	//	fmt.Println("this should not print2")
	//case (3 == 3):
	//	fmt.Println("print") //Without fallthrough case 3 prints only
	//	fallthrough
	//	case (4 == 4):
	//	fmt.Println("also true, and does print")
	//} //END - switch
	//* -------------------

	//No default fall through case. prints first true case only
	//switch  {
	//case false:
	//	fmt.Println("this should not print")
	//case (2 == 4):
	//	fmt.Println("this should not print2")
	//case (3 == 3):
	//	fmt.Println("print") //==> this case will only get printed
	//case (4 == 4):
	//	fmt.Println("also true, but should not print")
	//}  //END - switch
} //END - main

//DOCUMENTATION FOR switch
//There are two forms: expression switches and type switches
//the examples above are ALL expression switches

//In type switches the cases contain the types that are
//compared against the type of a specially annotated
//switch expression. (Conversion aka in C as casting)
//A missing switch expression is equivalent to a boolean value true
//default can occur anywhere in a switch statement
//the first case that equals the switch expression
//is executed--the other cases are skipped
//ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" . **The [] brackets mean OPTIONAL




