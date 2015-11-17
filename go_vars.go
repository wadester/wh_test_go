/*
 * Module:   go_vars.go
 * Purpose:  variable declarations in GO
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_vars.go
 * 2) Ref:  https://golang.org/pkg/fmt/
 *          http://www.dotnetperls.com/convert-go
*/

package main

// section for imports
import (
	"fmt"
	"math/cmplx"
)

// define vars -- note we can make our own section for vars!
var (
	fa,fb,fc = true, false, true  // auto type = bool
	s, x, y, z = "test", 1, 2, 3  // can be different types
	xx, yy, zz = 1+2i, 2+3i, 4+5i // complex uses i, not j like Python
	qq  complex128 = cmplx.Sqrt(1 + 1i)
)

// print them using %T for type of the value, and %v for value in default format
func main() {
	fmt.Printf("go_vars:  Simple types and their values using %%T(%%v)\n")
	const mypi = 3.1415
	const f = "%T(%v)\n"

	fmt.Printf(f, mypi, mypi)
	fmt.Printf(f, fa, fa)
	fmt.Printf(f, s, s)
	fmt.Printf(f, x, x)
	fmt.Printf(f, xx, xx)
	fmt.Printf(f, qq, qq)

	// short conversion, note explicit type conversion is REQUIRED
	fmt.Printf("Conversion tests\n")
	ii := 101
	dj := float64(ii)
	// this won't work:  qx := compex128(ii), nor qx := complex(ii, 0)
	// args to complex() must be float so 2 conversions required
	qx := complex(float64(ii), 0)
	re := real(qx)
	im := imag(qx)
	fmt.Printf(f, re, re)
	fmt.Printf(f, im, im)
	fmt.Printf(f, ii, ii)
	fmt.Printf(f, dj, dj)
	fmt.Printf(f, qx, qx)

	fmt.Printf("Done")
}

