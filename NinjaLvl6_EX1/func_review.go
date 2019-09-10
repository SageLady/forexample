package main

import "fmt"

func main()  {
	p := foo()
	fmt.Println(p)
	q, r := bar()
	fmt.Println("\n",q)
	fmt.Println("\n",r)
}

func foo() int {
	return 7
}

func bar() (i int, s string)  {

return 51,"A string"
}

//Hands-on exercise #1
//Review
//functions
//purpose of functions
//  **abstract code
//  **code reusability
//        ** DRY - Don’t Repeat Yourself
//func, receiver, identifier, params, returns
//func (receiver) identifier(parameters) (returns) { code }
//parameters vs arguments
//know the difference between parameters and arguments
//we define our func with parameters (if any)
//we call our func and pass in arguments (in any)
//Everything in Go is PASS BY VALUE
//variadic funcs
//Variadic parameter
// You can create a func which takes an unlimited number of arguments. When you do this, this is known as a “variadic parameter.” When use the lexical element operator “...T” to signify a variadic parameter (there “T” represents some type).

//Unfurling a slice
//When you have a slice of some type, you can pass in the individual values in a slice by using the “...” operator.

//multiple “variadic” params
//multiple “variadic” args
//returns
//multiple returns
//named returns - yuck!
//Defer
//A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
//personal anecdote: head down, ox plowing field; doing the work
//Methods
//A method is nothing more than a FUNC attached to a TYPE. When you attach a func to a type it is a method of that type. You attach a func to a type with a RECEIVER.
//Interfaces & polymorphism
//In Go, values can be of more than one type. An interface allows a value to be of more than one type. We create an interface using this syntax: “keyword identifier type” so for an interface it would be: “type human interface” We then define which method(s) a type must have to implement that interface. If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}), then that TYPE implicitly implements the interface and is also of that interface type. In Go, values can be of more than one type.

//Anonymous func
//Anonymous self-executing func

//func expressions
//Assigning a func to a variable

//Returning a func
//You can return a func from a func. Here is what that looks like. Returning a string or int. see samples
//callbacks
//passing a func as an argument
//functional programming not something that is recommended in go, however, it is good to be aware of callbacks
//idiomatic go: write clear, simple, readable code

//closure
//one scope enclosing another
//variables declared in the outer scope are accessible in inner scopes
//closure helps us limit the scope of variables

//recursion
//a func that calls itself
//factorial example

//life philosophy
//focus on what’s important; not upon what’s urgent

//Hands on exercise
//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

