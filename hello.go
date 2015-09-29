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

// factored input, recommended go style -- versus
//   import "fmt"
//   import "time
// note names are exported if they start with a capital letter
import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

// main function
func main() {
	fmt.Println("Welcome to GO!")
	
	fmt.Println("The time is", time.Now())
	rand.Seed(2)
	fmt.Println("Rand=",rand.Intn(10))
	fmt.Println("PI=", math.Pi)
}
