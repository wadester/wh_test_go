/*
 * Module:   go_method.go
 * Purpose:  demonstrate GO methods
 * Date:     10/07/2015
 * Notes:
 * 1) To build:
 *      go build go_method.go
 * 2) Ref:  https://tour.golang.org/methods/1
 * 3) GO does not have objects but does have methods on structs
 *    and these can be defined only on any times defined in my pgm.
 *
*/

package main

import (
	"fmt"
	"math"
)

// define a structure
type person struct {
	name, addr, city string
	zip  int
	age  int
	lat  float64
	lon  float64
}

// create a method on the structure to default to a location
// this is a pointer receiver and it is often used as it 
// prevents the copying data and it can modify the input
func (p *person) update_loc(lat, lon float64) {
	p.lat = lat
	p.lon = lon
}

// define a float type
type my_float float64

func (f my_float) to_int() int {
	return int(f)
}

func main() {
	p := &person{"Joe", "123 St", "Anytown", 123456, 21, 45.0, -77.0}
	fmt.Println("Person: ", p)
	p.update_loc(1.0, 2.0)
	fmt.Println("Person: ", p)

	f := my_float(math.Pi)
	i := f.to_int()
	fmt.Println("F=", f, "I=", i)
}
