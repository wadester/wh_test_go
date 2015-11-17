/*
 * Module:   go_bytes.go
 * Purpose:  very simple test of bytes 
 * Date:     N/A
 * Notes:
 * 1) To build:
 *      go build go_bytes.go
 * 2) Ref:  n/a
 *   
*/
package main

import (
        "fmt"
        "io"
        "bytes"
)

// simple function to write to the buffer
func ff(w io.Writer) {
         fmt.Fprintln(w, "This is a buffer write")
}

func main() {
	fmt.Println("go_bytes:  simple writer to a bytes buffer")
        var s1 = "This is a test string"
        fmt.Println("S1=", s1)

	bs := bytes.NewBufferString(s1)
	fmt.Println("BS=", bs.String())

        var b bytes.Buffer
        ff(&b)
        s2 := b.String()
        fmt.Println("S2=", s2)

}
