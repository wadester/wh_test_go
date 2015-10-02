/*
 * Module:   go_ptr.go
 * Purpose:  go pointers
 * Date:     10/2/2015
 * Notes:
 * 1) To build:
 *      go build go_ptr.go
 * 2) Ref:
 *
*/

package main

import (
	"fmt"
)

var (
	p1 *int
	i   int = 3
	pi  float64 = 3.14159
)

func main() {
	fmt.Println("go_ptr.go:  go pointer example")

	
	p2 := &i  // pointer to i, auto declared
	p1 = p2
	fmt.Printf("I=%d, *p1=%d, *p2=%d\n", i, *p1, *p2)
}
