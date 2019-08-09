package main

import "fmt"

func main(){
	s := "H"
	fmt.Println(s) //output: H
	bs := []byte(s)
	fmt.Println(bs) //output: [72] ascii capital H
	n := bs[0]
	fmt.Println(n)   //output: 72
	fmt.Printf("%T\n", n)  //type   //output: uint8
	fmt.Printf("%b\n", n)  //binary  //output: 1001000
	fmt.Printf("%#X", n)  //Hex    //0X48
}



//Numeral Systems:
//https://goo.gl/cvFUDa
//Output:
//H
//[72]
//72
//uint8
//1001000
//0X48
