/*
 * Module:   exercise-slices.go
 * Purpose:  Excecise for slices from GO tutorial
 * Date:     based on tutorial, N/A WWH mods
 * Notes:
 * 1) To build:
 *      go build exercise-slices.go
 * 2) Ref:  GO tutorial
 *    -- uncomment out the import and the pic.Show
 *    -- note use of two calls to "make()" to allocate the array
 */

package main


//import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	// need two makes to allocate the array, can't do dx*dy
	v := make([][]uint8, dy)
	for i := 0; i<dy; i++ {
		v[i] = make([]uint8, dx)
	}

	// make the pic
	for y:=0; y<dy; y++ {
		for x:=0; x<dx; x++ {
			v[y][x] = uint8(x*y)
		}
	}
	return(v)	
}
 
func main() {
	fmt.Println("exercise-slices:  Go slices exercise")

	// pic.Show(Pic)
	x := 128
	y := 256
	fmt.Println("Pic x=", x, "y=", y)
	v := Pic(y, x)
	fmt.Println("Len V=", v)
}
