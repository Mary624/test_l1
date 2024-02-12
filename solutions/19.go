package solutions

import (
	"fmt"
	"strings"
)

func Example19() {
	str := "Александр異"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	l := []rune(str)
	// Builder предоставляет эффективную конкатенацию строк
	var b strings.Builder
	b.Grow(len(str))
	for i := len(l) - 1; i >= 0; i-- {
		b.WriteRune(l[i])
	}
	return b.String()
}
