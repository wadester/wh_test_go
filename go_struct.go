/*
 * Module:   go_struct.go
 * Purpose:  go structures, arrays, and operators
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_struct.go
 * 2) Ref:  
 *
*/

package main

import (
	"fmt"
)

// define a 4d structure
type dim4 struct {
	X int
	Y int
	Z int
	t float64
}

// define some variables from our dim4
var (
	a dim4  // default values are 0
	b = dim4 {2, 3, 4, 3.1} // define defaults
	c = dim4 {X:2, t:9}    // selective assignement
	q [4]int       // array, fixed-size=4
	r [10]float32  // array, fixed-size=10
)

func main() {
	fmt.Println("go_struct:  go struct example")

	a.X = 1  // "." operator for member access
	a.Y = 2
	a.Z = 3
	a.t = 2.1

	d := dim4{3, 4, 5, 6.1} // auto assign w/ initialize
	p4 := &c  // auto assign pointer

	// create slice of ints
	s := []int{1,1,2,3,5,8,13}

	fmt.Println("A=", a)
	fmt.Println("B=", b)
	fmt.Println("C=", c)
	fmt.Println("D=", d)

	p4.X = 0  // pointer assignment uses "."
	fmt.Printf("C XYZ=%d %d %d\n", p4.X, p4.Y, p4.Z)

	// assign values to r, std array notation
	for i:=0; i<len(r); i++ {
		r[i] = float32(i)/2.0
	}
	fmt.Println("R=", r)
	fmt.Println("S=", s)

	ss := r[1:4]  // slice from r
	fmt.Println("SS=", ss)

	aa := make([]int, 5)  // use make() to build a slice
	fmt.Printf("%s len=%d cap=%d\n",
		aa, len(aa), cap(aa))
	// note:  The zero value of a slice is nil

	// use append to add to a slice, see notes
	var bb []int     // new nil slice
	bb = append(bb, 1) // add some values
	bb = append(bb, 2)
	bb = append(bb, 3)
	fmt.Println("Append example")
	for i := 0; i<len(bb); i++ {
		fmt.Printf("i=%d V=%d\n", i, bb[i])
	}
	// for with "range" format, replace "i" w/ _ to skip index
	fmt.Println("Range example")
	for i,v := range bb {
		fmt.Printf("i=%d V=%d\n", i, v)
	}
}
