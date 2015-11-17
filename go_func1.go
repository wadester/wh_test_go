/*
 * Module:   go_func1.go
 * Purpose:  second simple examples of functions in go
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_func1.go
*/
package main

import (
	"fmt"
)

// Test function returing two args 
// Note opening brace MUST be on same line as fcn due to auto 
// ";" insertion so the following will not build:
//      func swap(x, y string) (string, string) 
//      {
func swap(x, y string) (string, string) {
	return y, x
}

// Main function
func main() {
	fmt.Println("go_func1:  Go simple functions\n")

	var x string
	y := string("World") // short dec form
	x="Hello"
	fmt.Println("X=", x)
	fmt.Println("Y=", y)
	// Define a,b using short declaration form := 
	// Same as:
	//   var a string; var b string; a, b = swap(x,y);
	a, b := swap(x, y)
	fmt.Println("swap", a, b)
}
