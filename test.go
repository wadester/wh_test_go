package main

import (
	"fmt"
	"io"
	//"strings"
	"bytes"
)

func ff(w io.Writer) {
	 fmt.Fprintln(w, "Hello")
}

func main() {
	var s1 = "This is a test string"
	fmt.Println("S1=", s1)
	var b bytes.Buffer
	ff(&b)
	s2 := b.String()
	fmt.Println("S2=", s2)

}
