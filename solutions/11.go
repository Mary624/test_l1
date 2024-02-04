package solutions

import "fmt"

func Example11() {
	// у множеств уникальные числа
	l1 := []int{1, 55, 4, 2, 8, 100}
	l2 := []int{77, 4, 100, 5, 1}
	fmt.Println(findIntersection(l1, l2))
}

func findIntersection(l1, l2 []int) []int {
	// с помощью map проверяем, есть ли такие же числа во втором множестве
	m := make(map[int]bool, len(l1)+len(l2))
	res := make([]int, 0, len(l1)+len(l2))
	for _, v := range l1 {
		m[v] = true
	}
	for _, v := range l2 {
		_, ok := m[v]
		if ok {
			res = append(res, v)
		}
	}
	return res
}
