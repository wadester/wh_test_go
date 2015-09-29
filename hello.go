/*
 * Module:   hello.go
 * Purpose:  example hello world in GO
 * Date:     9/28/2015
 * Notes:
 * 1) To build:
 *      go build hello.go
 * 2) Ref:
 *      https://tour.golang.org/welcome
*/
package main

import (
	"fmt"
	"time"
)

// main function
func main() {
	fmt.Println("Welcome to GO!")

	fmt.Println("The time is", time.Now())
}
