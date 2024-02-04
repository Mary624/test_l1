package solutions

import "fmt"

func Example23() {
	l := []int{1, 2, 3, 4, 5}
	fmt.Println(deleteIndex(l, 4))
}

func deleteIndex(l []int, i int) []int {
	if i > len(l) || i < 0 {
		fmt.Println("index is out of range")
		return nil
	}
	return append(l[:i], l[i+1:]...)
}
