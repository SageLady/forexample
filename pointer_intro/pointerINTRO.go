package main

import "fmt"

func main()  {

	a := 42
	fmt.Println(a)   // OUTPUT: 42
	fmt.Println(&a) //  OUTPUT: Memory location, pointer returns memory location of `a`

	fmt.Printf("%T\n", a)  //Type that is `a`  OUTPUT: int
	fmt.Printf("%T\n", &a) //Type that is `&a`  OUTPUT: *int
	//Sharing an address:
	var b *int = &a
	fmt.Println(b)  // OUTPUT: the address of 'a' stored by `b`
    c := &a
    fmt.Println(*c)  //In the (*c) the `*` is an operator whereas the OUTPUT above *int the `*` is part of the type the OUTPUT to c := &a gives the value 42. (De-referencing the address means giving the value stored at an address)
    fmt.Println(*&a)   //OUTPUT: 42  (* with & means give the value stored at the address/memory location)

  *c = 43   // The value at memory location &c and &a is being changed to 43 from 42.
  fmt.Println(a)  //OUTPUT: 43
}

