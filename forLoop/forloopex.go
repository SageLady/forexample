package main

import "fmt"

func main()  {
	//Basic for loop:
	//for i := 0; i <= 2 ; i++  {
	//	fmt.Println(i)
	//}  //end for

	//nested loop example:
	//for j := 0; j <= 2; j++ {
	//	fmt.Println("Entering the inner loop:")
	//	for k := 0; k <= 1; k++  {
	//		fmt.Printf("The outer loop: %d\t The inner loop: %d\n", j, k)
	//	} //end inner for loop
	//} //end outer for loop

	//nested loop example 2:
	//for j := 0; j <= 2; j++ {
	//	fmt.Printf("The outer loop: %d\n", j)
	//	for k := 0; k <= 1; k++  {
	//		fmt.Printf("\t\tThe inner loop: %d\n", k)
	//	} //end inner for loop
	//} //end outer for loop

	//FOR SINGLE CONDITION example - SPECIFICATION BELOW
	//Runs as long as boolean condition resolves to True
	//How to set up a while statment
	//x := 1
	//for x < 5 {
	//		fmt.Println(x)
	//		x++
	//}
	//fmt.Println("done")
	//End first Single Condition example

    y := 1
	for {
			if y > 9 {
				break
			}  //end if
			fmt.Println(y)
			y++
	} //end for
	fmt.Println("done")
    //End second single Condition example

} //end main

//HIGHLIGHTS:
//The Go for loop is similar to--but not the same as C's.
//Unifies for and while
//No do-while loop
//three forms--only one uses semicolons

//For Single condition specification:
//For statements with single condition
//In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value true.

//for a < b {
//a *= 2
//}

//FOR STATEMENT WITH for CLAUSE
//A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement. The init statement may be a short variable declaration, but the post statement must not. Variables declared by the init statement are re-used in each iteration.

//ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .  [ ] brackets represent optional
//InitStmt = SimpleStmt .
//for i := 0; i < 10; i++ {
//f(i)
//}

//If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the semicolons are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value true.

//FOR STATEMENTS WITH RANGE CLAUSE
//For statements with range clause
//A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns iteration values to corresponding iteration variables if present and then executes the block.

//RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
//The expression on the right in the "range" clause is called the range expression, which may be an array, pointer to an array, slice, string, map, or channel permitting receive operations. As with an assignment, if present the operands on the left must be addressable or map index expressions; they denote the iteration variables. If the range expression is a channel, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the blank identifier, the range clause is equivalent to the same clause without that identifier.

//The range expression x is evaluated once before beginning the loop, with one exception: if at most one iteration variable is present and len(x) is constant, the range expression is not evaluated.

//Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

//Range expression                          1st value          2nd value

//array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
//string          s  string type            index    i  int    see below  rune
//map             m  map[K]V                key      k  K      m[k]       V
//channel         c  chan E, <-chan E       element  e  E
