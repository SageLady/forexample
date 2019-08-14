package main

import "fmt"

func main()  {
	//format printing example
	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n",true && false)
	fmt.Printf("true ||true\t %v\n",true || true)
	fmt.Printf("true || false\t %v\n",true || false)
	fmt.Printf("!true\t %v\n",!true)
	fmt.Printf("!false\t %v\n",!false)
}

//fmt.Println(true && true) //true AND true becomes true
//fmt.Println(true && false) //true AND false becomes false
//fmt.Println(true || true) //true OR true is true -- one or the other
//fmt.Println(true || false)  //true OR false becomes true
//fmt.Println(!true)  //not true becomes false

