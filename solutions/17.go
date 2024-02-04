package solutions

import "fmt"

func Example17() {
	l := []int{1, 2, 7, 10, 100, 231, 400}
	fmt.Println(binSearch(l, 400, 0))
}

func binSearch(l []int, v, left int) int {
	if len(l) == 1 {
		if l[0] == v {
			return left
		}
		fmt.Println("no such element")
		return -1
	}
	lL := len(l) / 2
	m := l[lL]
	if v == m {
		return lL + left
	}
	if v < m {
		return binSearch(l[:lL], v, left)
	}
	return binSearch(l[lL:], v, lL+left)
}
