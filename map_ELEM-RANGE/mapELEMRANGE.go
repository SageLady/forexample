package main

import "fmt"
//How to add something to map
//using for with range
func main()  {
	m := map[string]int{    //string is the key, and int is the value
		"James":32,     //key - james, int value of key = 32
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])  //grabbing something by using the key
	fmt.Println(m["something"])  //key doesn't exist and output is 0
	//adding a new key to the map
	m["todd"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}
	//slice of int
	//using for loop and range
	xi := []int{1,2,3,4,5}
	for i, v := range xi {
		fmt.Println(i,v)
	}

}
