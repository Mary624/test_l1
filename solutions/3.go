package solutions

import (
	"fmt"
	"sync"
)

func Example3() {
	l := []int{2, 4, 6, 8, 10}

	// WaitGroup позволяет дождаться выполнения всех горутин
	var wg sync.WaitGroup
	// для записи в общую переменную используем Mutex
	var m sync.Mutex
	var res int
	for _, v := range l {
		wg.Add(1)
		go square(v, &res, &wg, &m)
	}
	wg.Wait()
	fmt.Println(res)
}

func square(v int, res *int, wg *sync.WaitGroup, m *sync.Mutex) {
	v2 := v * v
	m.Lock()
	*res += v2
	m.Unlock()
	wg.Done()
}
