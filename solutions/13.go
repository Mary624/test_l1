package solutions

import "fmt"

func Example13() {
	var x = 3
	var y = 4
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	x, y = y, x
	fmt.Printf("After: x=%d, y=%d\n", x, y)
}
