package main

import "fmt"

func main()  {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
			}
	}
} //END main

//Hands-on exercise #7
//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."

//Range over the records, then range over the data in each record.
//solution: https://play.golang.org/p/FMM4c2PhZg
//video:  077

