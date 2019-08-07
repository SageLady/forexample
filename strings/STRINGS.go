package main

import "fmt"

func main(){
	s := "Hello Playground" //double quote
	bs := []byte(s) //Shown below: displays the string as ascii code
    fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")   //newline
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
//Blog by Rob Pike, one of the creators of Go
//https://blog.golang.org/strings
//UTF8  is the new ASCII
//RUNE is a code point in UTF8, which is represented by the  output
//U+0048 'H' U+0065 'e' U+006C 'l' U+006C 'l' U+006F 'o' U+0020 ' ' U+0050 'P' U+0
//06C 'l' U+0061 'a' U+0079 'y' U+0067 'g' U+0072 'r' U+006F 'o' U+0075 'u' U+006E
//'n' U+0064 'd'

//%s the uniterpreted bytes of the string or slice
//%q a double quoted string safely escaped with Go syntax
//%x base 16, lower-case, two chars per byte
//%X base 16, upper-case, two chars per byte
//Alternate format: add leading 0 for octal (%#o)
//0x for hex (%#x), 0X for hex (%#X), suppress
//0x for %p (%#p), for %q, print a raw (backquoted) string
//if strconv.CanBackquote returns true
//OUTPUT:
//%#U
//U+0048 'H' U+0065 'e' U+006C 'l' U+006C 'l' U+006F 'o' U+0020 ' ' U+0050 'P' U+0
//06C 'l' U+0061 'a' U+0079 'y' U+0067 'g' U+0072 'r' U+006F 'o' U+0075 'u' U+006E
//'n' U+0064 'd'

