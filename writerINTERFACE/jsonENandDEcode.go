package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("This function is calling Fprintln shown below")
	fmt.Fprintln(os.Stdout, "Using os.Stdout in this and below")

	io.WriteString(os.Stdout, "Hello from WriteString")
}
//https://golang.org/pkg/io/#Writer
//type Writer ¶
//Writer is the interface that wraps the basic Write method.

//Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.

//Implementations must not retain p.

//type Writer interface {
//	Write(p []byte) (n int, err error)
//}
//-------------------------------------------------

//Writer interface
//Understanding the writer interface from package io. Also, one last note about working with JSON: encode & decode.
//code: https://play.golang.org/p/3Txh-dKQBf
//video: 121

//CLASS PROGRAM:
//package main

//import (
//"fmt"
//"io"
//"os"
//)

//func main() {
//	fmt.Println("Hello, playground")
//	fmt.Fprintln(os.Stdout, "Hello, playground")
//	io.WriteString(os.Stdout, "Hello, playground")
//}

//https://golang.org/pkg/encoding/json/
//Notes: Decoder and encoder work without the
//need for variables to help translate/convert
//json to go and go to json. Gets data from
//the wire and does the process by using decoder or
//encoder. The processes are either using a reader or
//writer. Decoder = Reader and Encoder = Writer
//The reader and writer are doing the action from sources
//that include but not limited to: a file, a web page or
//connection
//type Decoder  (liken to marshal)
//func NewDecoder(r io.Reader) *Decoder
//func (dec *Decoder) Buffered() io.Reader
//func (dec *Decoder) Decode(v interface{}) error
//func (dec *Decoder) DisallowUnknownFields()
//func (dec *Decoder) More() bool
//func (dec *Decoder) Token() (Token, error)
//func (dec *Decoder) UseNumber()
//type Delim
//func (d Delim) String() string
//type Encoder  (liken to unmarshal)
//func NewEncoder(w io.Writer) *Encoder
//func (enc *Encoder) Encode(v interface{}) error
//func (enc *Encoder) SetEscapeHTML(on bool)
//func (enc *Encoder) SetIndent(prefix, indent string)