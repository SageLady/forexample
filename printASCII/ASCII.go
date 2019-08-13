package main
//loop over numbers 33 - 122
//ASCII upper case:  65 through 90
//ASCII lower case:
//print them out as text 97 through 122
//print them out as a number
//and a string
//Integer (fmt library https://godoc.org/fmt)
//base 10 which equals decimal ==> %d
//base 16, with lower case letters for a -f ==> %x
//base 16, which upper case letters for A - F ==> %X
//Unicode format: U+1234; same as "U+%04X"  ==> %U
//The value of integer when format string
//is being used is %v
//hexadecimal = %#x
//Unicode = %#U
//Hex and Unicode coincide at the unicode code points, the
// 0x74    U+0074 't'  (The 74)
import "fmt"

	func main()  {
		for i := 33; i <= 122; i++ {
			fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
		}
	}
