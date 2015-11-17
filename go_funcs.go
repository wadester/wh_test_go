/*
 * Module:   go_funcs.go
 * Purpose:  go functional programming misc
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_funcs.go
 * 2) Ref:
 *
*/

package main

import (
	"fmt"
	"math"
)

// closure example, from examples at:
//   https://tour.golang.org/moretypes/21
// adder is a func with no args and it returns a 
// function that returns and int:  "func(int) int"
func adder() func(int, int) int {
	sum := 0
	idx := 0
	return func(i,x int) int {
		sum += x + idx
		idx = i
		return sum
	}
}

func main() {
	fmt.Println("go_funcs:  test of Go functional programming")

	
	fmt.Println("Test of variable equal to a func")
	sqr := func(x float64) float64 {
		return(x*x)
	}

	x := math.Pi
	xx:=sqr(x)
	fmt.Println("X=", x, "XX=", xx)

	// create two "adders" and call them, each accmulates
	fmt.Println("Create closures and use them")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i,i),
			neg(i, -2*i),
		)
	}


}
