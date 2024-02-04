package solutions

import "fmt"

func Example10() {
	l := []float64{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
	// с помощью map проверяем, есть ли такой диапазон
	m := make(map[int][]float64, len(l))
	for _, v := range l {
		r := int(v) / 10
		r *= 10
		_, ok := m[r]
		if !ok {
			m[r] = make([]float64, 0, len(l))
		}
		m[r] = append(m[r], v)
	}
	fmt.Println(m)
}
