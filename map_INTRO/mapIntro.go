package main

import "fmt"

func main()  {
	m := map[string]int{    //string is the key, and int is the value
		"James":32,     //key - james, int value of key = 32
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])  //grabbing something by using the key
	fmt.Println(m["something"])  //key doesn't exist and output is 0
	//How to check the map for a key to see if it exists
	//comma okay idiom
	v, ok  := m["something"]
	fmt.Println(v)
	fmt.Println(ok)
	//Can convert the statements into conditional statements
	if v, ok := m["something"]; ok {
		fmt.Println("This is the IF print",v)
	} else {  //END if (if false then do something
		fmt.Println("DOES NOT EXIST", v)
	}
	//adding the



}
