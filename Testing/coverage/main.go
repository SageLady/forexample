package main
//godoc.org/testing
import (
	"fmt"
	"github.com/GoesToEleven/forexample1/Testing/benchmarkExamples/03-cat/mystr"
	"strings"
)

const s = "VERY COOL STATEMENT BY MARIANNE WILLIAMSON"

//var xs []string

func main()  {
	xs := strings.Split(s, " ")

	for  _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}

//Set up:
//**#1: 02-greet - program file contains main.go
//mystr (my string package) - contained in 02-greet directory, is new folder.
//mystr package folder - contains main.go and main_test.go
//** #2: mystr package-->main.go - small program
//** #3: mystr package-->main_test.go - contains functions named using Test, Example and Benchmark. Like TestGreet, ExampleGreet and BenchmarkGreet

// -----------------------------------------------------
//** #1: START - 02-greet - main program
//package main

//import (
//"fmt"
//"github.com/GoesToEleven/go-programming/code_samples/009-testing/04-benchmark/02-greet/mystr"
//)

//func main() {
//	fmt.Println(mystr.Greet("James"))
//}
//** #1: END - 02-greet - main program
// -----------------------------------------------------
//** #2: START - mystr package-->main.go
//package mystr

//import "fmt"

//func Greet(s string) string {
//	return fmt.Sprint("Hello my dear, ", s)
//}
//** #2: END - mystr package-->main.go
// -----------------------------------------------------
//** #3: START - mystr package-->main_test.go
//package mystr

//import (
//"fmt"
//"testing"
//)

//func TestGreet(t *testing.T) {
//	s := Greet("James")
//	if s != "Hello my dear, James" {
//		t.Error("got", s, "want", "Hello my dear, James")
//	}
//}

//func ExampleGreet() {
//	fmt.Println(Greet("James"))
	// Output:
	// Hello my dear, James
//}

//func BenchmarkGreet(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Greet("James")
//	}
//}
//** #3: END - mystr package-->main_test.go
// -----------------------------------------------------

//Coverage
//Coverage in programming is how much of our code is covered by tests. We can use the “-cover” flag to run coverage analysis on our code. We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.
//code:
//go test -cover
//go test -coverprofile c.out
//show in browser:
//go tool cover -html=c.out
//learn more
//go tool cover -h
//https://github.com/GoesToEleven/go-programming
//video: 186
//Remember to BET
//Benchmark
//Example
//Test

//BenchmarkCat(b *testing.B)
//ExampleCat()
//TestCat(t *testing.T)

//Commands -command line
//go test
//go test -bench ==> can be used with regular expression, -bench=X/Y
//go test -cover
//go test -coverprofile c.out
//go tool cover -html=c.out

// -- go test --
//go test [build/test flags] [build/test flags & test binary flags]
//Automates testing the packages named by the import paths
//Prints summary of the test results
//Re-compiles each package along with any fils with name matching
//the file pattern "*_test.go"
//Ignores directory named "testdata," making it available
//to hold ancillary data needed ty the tests.
//Runs in two different modes:
//1. local directory mode - used when go test is invoked
//with no package arguments (for example, 'go test' or 'go test -v).
//2. Packe list mode - used when go test is invoked  with
//explicit package arguments ( for example 'go test math', or 'go
// test ./...' and even 'go test .' In this mode, go test compiles
//and tests each of the packages listed in the command line. If
//invoked with the -bench or -v flag, go test prints the full output even for passing package tests, in order to display the reequested
//benchmark result or verbose logging.

// -bench = To run all benchmarks, use '-bench' or '-bench=.'
// -benchtime t = run enough iterations of each benchmark to take t, specified as a time. The default is 1 second (1s)
// -count n = run each test and benchmark n times (default 1).
// -cover = enables coverage analysis. Coverage works by annotating the source code before compilation, compilation and test failures with coverage enabled may report line numbers that don't correspond to the original sources.
// -coverprofile cover.out = write a coverage profile to the file after all tests have passed
// -trace trace.out = write an execution trace to the  specified file before exiting
