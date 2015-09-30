/*
 * Module:   go_vars.go
 * Purpose:  variable declarations
 * Date:     9/20/2015
 * Notes:
 * 1) To build:
 *      bo build go_vars.go
 * 2) Ref:  https://golang.org/pkg/fmt/
 *
*/

package main

import (
	"fmt"
	"math/cmplx"
)

// define vars
var (
	fa,fb,fc = true, false, true  // auto type = bool
	s, x, y, z = "test", 1, 2, 3  // can be different types
	xx, yy, zz = 1+2i, 2+3i, 4+5i // complex uses i, not j like Python
	qq  complex128 = cmplx.Sqrt(1 + 1i)
)

// print them using %T for type of the value, and %v for value in default format
func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, fa, fa)
	fmt.Printf(f, s, s)
	fmt.Printf(f, x, x)
	fmt.Printf(f, xx, xx)
	fmt.Printf(f, qq, qq)
	fmt.Printf("Done")
}

