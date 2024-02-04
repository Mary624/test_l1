package solutions

import (
	"strings"
)

func Example15() {
	someFunc()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	v = v[:100]
	c := make([]rune, len(v))
	// без copy сохранится ссылка на огромную строку
	copy(c, []rune(v))
	justString = string(c)
}

func createHugeString(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteString("a")
	}
	return b.String()
}
