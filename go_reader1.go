/*
 * Module:   go_reader1.go
 * Purpose:  basic IO reader based on examples in tutorial
 * Date:     ?, N/A WWH mods
 * Notes:
 * 1) To build:
 *      go build go_reader1.go
 * 2) Ref:
 *
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Printf("go_reader1:  Go Reader example from tutorial\n")

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
