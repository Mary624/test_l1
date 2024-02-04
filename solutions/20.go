package solutions

import (
	"fmt"
	"strings"
)

func Example20() {
	str := "Александр  異 кот Жизнь snow "
	fmt.Println(reverseWords(str))
}

func reverseWords(str string) string {
	l := []rune(str)
	// Builder предоставляет эффективную конкатенацию строк
	var b strings.Builder
	var bW strings.Builder
	for i := len(l) - 1; i >= 0; i-- {
		if l[i] != ' ' || i == 0 {
			bW.WriteRune(l[i])
			continue
		}
		if bW.Len() == 0 {
			continue
		}
		b.WriteString(reverse(bW.String()) + " ")
		bW.Reset()

	}
	b.WriteString(reverse(bW.String()))
	return b.String()
}
