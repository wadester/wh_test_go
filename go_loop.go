/*
 * Module:   go_loop.go
 * Purpose:  example GO loops 
 * Date:     9/30/2015
 * Notes:
 * 1) To build:
 *      go build go_loop.go
 * 2) Ref:
 *
*/

package main
import "fmt"

func main() {

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

}

