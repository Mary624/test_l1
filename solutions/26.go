package solutions

import (
	"fmt"
	"strings"
)

func Example26() {
	fmt.Println(checkRepeats("abcd"))
	fmt.Println(checkRepeats("abCdefAaf"))
	fmt.Println(checkRepeats("aabcd"))
}

func checkRepeats(str string) bool {
	// проверяем уникальность с помощью map
	m := make(map[string]bool, len(str))
	for _, r := range str {
		l := strings.ToLower(string(r))
		_, ok := m[l]
		if ok {
			return false
		}
		m[l] = true
	}
	return true
}
