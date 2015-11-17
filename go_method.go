/*
 * Module:   go_method.go
 * Purpose:  demonstrate Go methods 
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_method.go
 * 2) Ref:  https://tour.golang.org/methods/1
 * 3) Go does not have objects but does have methods on structs
 * 4) Go receivers can be pointer or value.  Pointer receivers avoid
 *    copying and can change the receiving structure.
 *     
 *
*/

package main

import (
	"fmt"
	"math"
)

// Interface for the Person structure.
// Can be used with anything that implements these methods....
type Aperson interface {
	Update_Name(string)
	Update_Loc(lat, lon float64)
}

// Simple interface implementing ToInt only.
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

// Create a method on the structure to default to a location.
// This is a pointer receiver and it is often used as it 
// prevents copying data and it can modify the input.
func (p *Person) Update_Loc(lat, lon float64) {
	p.lat = lat
	p.lon = lon
}

// Second method, again using pointer as we update the struct.
func (p *Person) Update_Name(n string) {
	p.name = n
}

// Define a stringer for the Person.
// This defines how object is presented as a string.
// This can be by value as we don't change the struct.
func (p Person) String() string {
	return fmt.Sprintf("%v %v %v %v %v years (%v, %v)",
		p.name, p.addr, p.city, p.zip, p.age, p.lat, p.lon)
}

// Define ToInt for the Person structure
func (p Person) ToInt() int {
	return p.age
}

// Define a float type and a method on it, in this case make it an int
type my_float float64

func (f my_float) ToInt() int {
	return int(f)
}

// main function:
func main() {
	fmt.Println("go_method:  Go methods example")

	p := &Person{"Joe", "123 St", "Anytown", 123456, 21, 45.0, -77.0}
	fmt.Println("Person: ", p)
	p.Update_Loc(1.0, 2.0)
	fmt.Println("Person: ", p)

	f := my_float(math.Pi)
	i := f.ToInt()
	fmt.Println("F=", f, "I=", i)

	// use the first interface
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
