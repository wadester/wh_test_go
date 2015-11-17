/*
 * Module:   go_reader2.go
 * Purpose:  second reader example
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_reader2.go
 * 2) Ref:  n/a
 * 3) This uses buffers and writes to buffers
 *   
*/
package main

import (
	"fmt"
	"io"
	//"strings"
	"bytes"
)

// make writer function to write to its output
func fw(w io.Writer) {
	fmt.Fprintln(w, "Writer Function Wrote This")
}

// make writer function that copies input
func fc(s string) *bytes.Buffer{
	b := bytes.NewBufferString(s)
	return b
}

func main() {
	fmt.Println("go_reader2:  Go reader 2 example")

	// make buffer and use writer to write to it, then
	// convert it to a string
	var bb bytes.Buffer
	fw(&bb)
	s0 := bb.String() 
	fmt.Println("S0=", s0)

	// make and print input string
	s := "This is a test string that will be accessed"
	sl := len(s)
	fmt.Println("S (len=", sl, ") = ", s)
	
	// make buffer from string
	// -- not as efficient as strings.NewReader()
	bs := bytes.NewBufferString(s)
	fmt.Println("BS=", bs.String())

	// func call to make buffer from string
	ba := fc(s)
	fmt.Println("BA=", ba.String())

	// use the read method on the buffer
	b := make([]byte, 8)
	for {
		n, err := ba.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
