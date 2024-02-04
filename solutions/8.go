package solutions

import (
	"fmt"
	"math"
)

func Example8() {
	var v int64 = 5
	i := 1
	// меняет i-ый элемент начиная с 0
	fmt.Println(v ^ step2(i))
}

func step2(i int) int64 {
	return int64(math.Pow(2, float64(i)))
}
