/*
 * Module:   go_func.go
 * Purpose:  simple examples of functions in go
 * Date:     9/29/2015
 * Notes:
 * 1) To build:
 *      go build go_func.go
*/
package main

import (
	"fmt"
	"time"
)

func ptime() {
	fmt.Println("The time is", time.Now())
}

func add(x int, y int) int {
	return x + y
}

// main function
func main() {
	var x int
	var y int
	x=100
	y=101
	fmt.Println("X=", x)
	fmt.Println("Y=", y)
	fmt.Println("X+Y=", add(x,y))
	ptime()
}
