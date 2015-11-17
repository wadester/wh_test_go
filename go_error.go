/*
 * Module:   go_error.go
 * Purpose:  GO Error processing test
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_error.go
 * 2) Ref:  https://tour.golang.org/methods/8
 *
*/

package main
import (
	"fmt"
	"time"
	"strconv"
)

// Error structure with time and a message string
type MyError struct {
	When time.Time
	What string
}

// Implement the Error() function, implementing the interface
// format the string with the time and message
//      type error interface {
//          Error() string
//      }
//
func (e *MyError) Error() string {
	return fmt.Sprintf("ERROR:  %v, %s", e.When, e.What)
}

// Dummy function called run() to return an error w/ time and msg
func run() error {
	return &MyError{
		time.Now(),
		"run() failed",
	}
}

// Main function, calls the run() and checks for errors
func main() {
	fmt.Println("go_error.go:  Error examples")
	fmt.Println("Test dummy function returning error (should report error)")

	// precursor   ; if test   -- if with initialization
	if err := run(); err != nil {
		fmt.Println(err)
	}

	istr:="42"
	fmt.Println("Trying to convert", istr, " (should be OK)")
	i, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", i)

	istr="xx"
	fmt.Println("Trying to convert", istr, " (should fail)")
	i, err = strconv.Atoi(istr)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", i)

}
