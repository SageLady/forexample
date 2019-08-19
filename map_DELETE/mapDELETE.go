package main

import "fmt"

func main()  {
	m := map[string]int{    //string is the key, and int is the value
		"James":32,     //key - james, int value of key = 32
		"Miss Moneypenny":27,
		"Q":80,
		"R":45,
		"S":55,
		"T":78,
		"U":66,
	}
	fmt.Println(m)
	//deleting from map format: delete(<map name>, "key")
	delete(m,"S")
	fmt.Println(m)
	//If a key does not exist no error is thrown.
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value", v)
		delete(m, "Miss Moneypenny")
	}
    fmt.Println(m)
}
