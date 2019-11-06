package main

import (
	"fmt"
	"github.com/GoesToEleven/forexample1/NinjaLvl12/01/dog"
)

type canine struct {
	name string
	age int
}

func main()  {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}

//Hands-on exercise #1
//Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years). Document  your code with comments. Use this code in func main.
//run your program and make sure it works
//run a local server with godoc and look at your documentation.
//solution: https://github.com/GoesToEleven/go-programming
//video: 180-f
///Hands-on exercise #2
//Push the code to github. Get your documentation on godoc.org and take a screenshot. Delete your code from github. Refresh godoc.org so that it no longer has your code. Tweet me (https://twitter.com/Todd_McLeod) your screenshot.
//solution: https://github.com/GoesToEleven/go-programming
//video: no video
//Hands-on exercise #3
//Use godoc at the command line to look at the documentation for:
//fmt
//fmt Print
//strings
//strconv
//solution: https://github.com/GoesToEleven/go-programming
//video: no video
