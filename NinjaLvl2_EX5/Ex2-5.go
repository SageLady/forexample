package main

import "fmt"

func main()  {
t := `"Hello World, 
HERE I COME"`
fmt.Println(t)
}

//Hands on excercise #5
//Create a variable of type string using a
//raw string literal. Print it
//Solution: https://play.golang.org/p/dLy36A-V-w

//Definition:
//String literals ¶
//A string literal represents a string constant obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

//Raw string literals are character sequences between back quotes, as in `foo`. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.