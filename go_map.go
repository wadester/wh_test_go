/*
 * Module:   go_map.go
 * Purpose:  go maps
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_map.go
 * 2) Ref:
 *
*/

package main

import (
	"fmt"
)

// define a 4D structure
type dim4 struct {
	X, Y, Z int
	t float64
}

var m1 map[int]dim4
var m2 map[string]dim4
var m3 = map[string]dim4{
	"Altair": dim4{1,2,3,4.0},  // trailing comma is needed!
}

func main() {
	fmt.Println("go_map:  Go example of maps")

	m1 = make(map[int]dim4)  // must make first
	m1[1234] = dim4{1,2,3,4.0}
	m1[1234] = dim4{2,3,4,5.0}
	m1[666] = dim4{-1,-1,-1,2.7}
	fmt.Println("M1=", m1)
	m2 = make(map[string]dim4)  // must make first
	m2["Polaris"] = dim4{1,2,3,4.0}
	m2["Vega"] = dim4{1,2,3,4.0}
	fmt.Println("M2=", m2)
	fmt.Println("M3=", m3)

	s := "Polaris"
	e := m2[s]
	fmt.Println("S=", s, "E=", e)

	// one-line test for key:  https://tour.golang.org/moretypes/18
	s = "Sirius"
	ee, ok := m2[s]
	fmt.Println("Key=", s, "v=", ee, "Present?", ok)
}
