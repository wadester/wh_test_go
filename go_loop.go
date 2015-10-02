/*
 * Module:   go_loop.go
 * Purpose:  example GO loops, if, switch, etc. (flow control)
 * Date:     9/30/2015
 * Notes:
 * 1) To build:
 *      go build go_loop.go
 * 2) Ref:  GO tutorial on flowcontrol
*           https://tour.golang.org/flowcontrol
 *
*/

package main
import "fmt"

// test function, called when main exits
func MyFunc(str string) {
	fmt.Println(str)
}

// module main function
func main() {

	// defered function, called when main exits
	defer MyFunc("Deferred Bye call, goes on LIFO stack")

	fmt.Println("go_loop.go:  GO loops and loop notes\n");

	// only loop type is for, parens not allowed, must use {}
	var ii int
	for ii=0; ii<3; ii++ {
		fmt.Printf("II=%d\n", ii)
	}

	// using := will auto-define i below
	for i:=0; i<3; i++ {
		fmt.Printf("i=%d\n", i)
	}

	// for used like a while, note ;'s can be removed
	// also note that inf loop can be written as "for {"
	ii=0
	for ii<10 {
		ii++

		// typical if statement, can include optional "(" 
		if (ii==2) {
			fmt.Println("We found 2")
		}
		if ii==3 {
			fmt.Println("We found 3")
		}

		// if can include a pre-statement
		if jj:=ii*2; ii==4 {
			fmt.Printf("jj=%d\n", jj)
		}
	}

	/* switch like C but implied "break" */
	for xx:=1; xx<11; xx++ {
		switch yy := xx*xx; yy {
		case 1:
			fmt.Println("One!")
		case 4:
			fmt.Println("Two squared")
		case 9:
			fmt.Println("Three squared and")
			fallthrough  // opposite of C's break....
		default:
			fmt.Printf("Many=%d,%d\n", xx, yy)
		}
	}

	xx:=1
	//yy:=2
	// switch w/o condition is like switch true
	// this is useful like multiple IF statements
	switch {
	case xx==1:
		fmt.Println("XX is one!")
	default:
		fmt.Println("Should not get here")
	}
}

