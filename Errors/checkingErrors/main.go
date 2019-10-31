package checkingErrors

//Checking errors
//Write the code with errors before writing the code without errors. Always check for errors. Always, always, always.*
//(*almost always)
//code:
//https://github.com/GoesToEleven/go-programming
//video: 172


//CODE SAMPLE ONE:
//package main

//import "fmt"

//func main() {
//	n, err := fmt.Println("Hello")
//	if err != nil {        //If err equals nil there is no error. If err != nil then error.
//		fmt.Println(err)
//	}
//	fmt.Println(n)   //OUTPUT: prints out the number of bytes, the newline character counts
//}


//CODE SAMPLE TWO:
//This code uses the scan method to receive user input. Code can be
//run from command line to test.
//package main

//import "fmt"

//func main() {
//	var answer1, answer2, answer3 string

//	fmt.Print("Name: ")
//	_, err := fmt.Scan(&answer1)   //Scan for input,
//	if err != nil {
//		panic(err)
//	}

//	fmt.Print("Fav Food: ")
//	_, err = fmt.Scan(&answer2)
//	if err != nil {
//		panic(err)
//	}

//	fmt.Print("Fav Sport: ")
//	_, err = fmt.Scan(&answer3)
//	if err != nil {
//		panic(err)
//	}

//	fmt.Println(answer1, answer2, answer3)

//}


//CODE SAMPLE THREE:
//package main

//import (
//"fmt"
//"io"
//"os"
//"strings"
//)

//func main() {
//	f, err := os.Create("names.txt")  //This creates a text file named, "names.txt"
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()   //defer closing the file

//	r := strings.NewReader("James Bond")    //string is a reader string

//	io.Copy(f, r)  //the reader string goes to the writer, 'f'
//}

//CODE SAMPLE FOUR:
//package main

//import (
//"fmt"
//"io/ioutil"
//"os"
//)

//Code will open a file, read the file, and display the text in the file

//func main() {
//	f, err := os.Open("names.txt")   //Opens file, "names.txt"
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()    //defer closing file.

//	bs, err := ioutil.ReadAll(f)    //takes in a reader and returns a byte slice and error
//	if err != nil {                 //
//		fmt.Println(err)
//		return
//	}

//	fmt.Println(string(bs))
//}