/*
 * Module:   go_ptr.go
 * Purpose:  go pointers
 * Date:     N/A
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
	fmt.Println("go_ptr:  Go pointer example")

	p2 := &i  // pointer to i, auto declared
	p1 = p2
	fmt.Printf("I=%d, *p1=%d, *p2=%d\n", i, *p1, *p2)
}
