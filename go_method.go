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

// interface for the Person structure
// can be used with anything that implements these methods....
type Aperson interface {
	Update_Name(string)
	Update_Loc(lat, lon float64)
}

// simple interface implementing ToInt only
type Ati interface {
	ToInt() int
}

// define a structure
type Person struct {
	name, addr, city string
	zip  int
	age  int
	lat  float64
	lon  float64
}

// create a method on the structure to default to a location
// this is a pointer receiver and it is often used as it 
// prevents the copying data and it can modify the input
func (p *Person) Update_Loc(lat, lon float64) {
	p.lat = lat
	p.lon = lon
}

// second method
func (p *Person) Update_Name(n string) {
	p.name = n
}

// define a stringer for the Person -- how object is presented as a string
func (p Person) String() string {
	return fmt.Sprintf("%v %v %v %v %v years (%v, %v)",
		p.name, p.addr, p.city, p.zip, p.age, p.lat, p.lon)
}

// define ToInt for the Person structure
func (p Person) ToInt() int {
	return p.age
}

// define a float type and a method on it
type my_float float64

func (f my_float) ToInt() int {
	return int(f)
}

func main() {
	p := &Person{"Joe", "123 St", "Anytown", 123456, 21, 45.0, -77.0}
	fmt.Println("Person: ", p)
	p.Update_Loc(1.0, 2.0)
	fmt.Println("Person: ", p)

	f := my_float(math.Pi)
	i := f.ToInt()
	fmt.Println("F=", f, "I=", i)

	// use the interface
	var a Aperson
	a = p
	fmt.Println("A=", a)


	// use the second interface
	var b Ati
	b = p
	fmt.Println("B assigned from p=", b, b.ToInt())
	b = f
	fmt.Println("B assigned from f=", b, b.ToInt())

}
