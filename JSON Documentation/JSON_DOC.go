package main

func main()  {
	
}
//www.godoc.org/encoding/json
//JSON
//ENCODING --> JSON (Level down to JSON within Go-->Documentation)
//package json
//EXAMPLE: (with notes on documentation)
//type ColorGroup struct {
//	ID     int     // WHEN CAPITALIZED means variable is
//	Name   string  // accessible from the outside
//	Colors []string  // touching upon private and public
//}
//group := ColorGroup{
//ID:     1,
//Name:   "Reds",
//Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
//}
//b, err := json.Marshal(group)
//if err != nil {
//fmt.Println("error:", err)
//}
//os.Stdout.Write(b)

//POINTERS HELP US USE METHODS:
//For example:
//If you have a pointer to Decoder, then you have access
//to all the methods associated to Decoder. Shown below

//type Decoder   (type pointer to Decoder

//Methods for Decoder
// (r io.Reader = parameter - passing in reader),
//  func NewDecoder returns a pointer to Decoder
//func NewDecoder(r io.Reader) *Decoder   //create
//func (dec *Decoder) Buffered() io.Reader
//func (dec *Decoder) Decode(v interface{}) error
//func (dec *Decoder) DisallowUnknownFields()
//func (dec *Decoder) More() bool
//func (dec *Decoder) Token() (Token, error)
//func (dec *Decoder) UseNumber()

//If you have a pointer to Delim you have access
//to the methods associated with Delim

//type Delim
//func (d Delim) String() string

//If you have a pointer to Encoder you have access
//to the methods associated with Encoder

//type Encoder
//func NewEncoder(w io.Writer) *Encoder
//func (enc *Encoder) Encode(v interface{}) error
//func (enc *Encoder) SetEscapeHTML(on bool)
//func (enc *Encoder) SetIndent(prefix, indent string)


