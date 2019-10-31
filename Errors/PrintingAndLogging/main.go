package PrintingAndLogging

import (
	"fmt"
	"log"
	"os"
)

//Printing & logging
//You have a few options to choose from when it comes to printing out, or logging, an error message:
//fmt.Println()    \\Goes to standard IO
//log.Println()    \\Goes to standard IO, but can also go to a file
//log.Fatalln()    \\Fatal line - exit code  Deferred functions do not run.
//os.Exit()
//log.Panicln()   \\Panic - is a panic or warning, but code continues. Deferred functions run.
//deferred functions run
//can use “recover”
//panic()
//code:
//https://github.com/GoesToEleven/go-programming
//video: 173

//package main

//import (
//"fmt"
//"os"
//)

//func main() {
//	_, err := os.Open("no-file.txt")
//	if err != nil {
//		fmt.Println("err happened", err)   //Customize error with unique text
		//		log.Println("err happened", err)   //log also gives us a date and timestamp
		//		log.Fatalln(err)
		//		panic(err)
//	}
//}

// Println formats using the default formats for its operands and writes to standard output.


//SAMPLE CODE TWO:
//package main

//import (
//"fmt"
//"log"
//"os"
//)

//func main() {
//	f, err := os.Create("log.txt")   //IF error is generated then a text file will be created
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer f.Close()
//	log.SetOutput(f)

//	f2, err := os.Open("no-file.txt")
//	if err != nil {
//		//		fmt.Println("err happened", err)
//		log.Println("err happened", err)
//		//		log.Fatalln(err)
//		//		panic(err)
//	}
//	defer f2.Close()

//	fmt.Println("check the log.txt file in the directory")
//}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

//SAMPLE CODE THREE: demonstration of fatal error
//package main

//import (
//"fmt"
//"log"
//"os"
//)

//func main() {
//	defer foo()
//	_, err := os.Open("no-file.txt")
//	if err != nil {
//		//		fmt.Println("err happened", err)
//		//		log.Println("err happened", err)
//		log.Fatalln(err)
//		//		panic(err)
//	}
//}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message ...
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
