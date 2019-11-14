package main

import "testing"

//Should be in the same package

func TestMySum(t *testing.T)  {   //MySum is capitalized. Can be called whatever, best practice.
	x := mySum(2,3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}

}
