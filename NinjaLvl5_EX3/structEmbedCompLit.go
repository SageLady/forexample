package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}



func main()  {
t1 := truck{
	vehicle:   vehicle{
		doors: 4,
		color: "blue",
	},
	fourwheel: true,
}

	s1 := sedan{
		vehicle:   vehicle{
			doors: 2,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Printf("Truck color is: %v", t1.color)
	fmt.Printf("\nSedan color is: %v", s1.color)

} //END main
//Hands-on exercise #3
//Create a new type: vehicle.
//The underlying type is a struct.
//The fields:
//doors
//color
//Create two new types: truck & sedan.
//The underlying type of each of these new types is a struct.
//Embed the “vehicle” type in both truck & sedan.
//Give truck the field “fourWheel” which will be set to bool.
//Give sedan the field “luxury” which will be set to bool. solution
//Using the vehicle, truck, and sedan structs:
//using a composite literal, create a value of type truck and assign values to the fields;
//using a composite literal, create a value of type sedan and assign values to the fields.
//Print out each of these values.
//Print out a single field from each of these values.
//solution: https://play.golang.org/p/PrTtTv_vVO
//video: 088

//SOLUTION:
//package main

//import (
//"fmt"
//)

//type vehicle struct {
//	doors int
//	color string
//}

//type truck struct {
//	vehicle
//	fourWheel bool
//}

//type sedan struct {
//	vehicle
//	luxury bool
//}

//func main() {
//	t := truck{
//		vehicle: vehicle{
//			doors: 2,
//			color: "white",
//		},
//		fourWheel: true,
//	}

//	s := sedan{
//		vehicle: vehicle{
//			doors: 4,
//			color: "black",
//		},
//		luxury: false,
//	}
//	fmt.Println(t)
//	fmt.Println(s)
//	fmt.Println(t.doors)
//	fmt.Println(s.doors)
//}