package solutions

import (
	"fmt"
	"sync"
)

func Example2() {
	l := []int{2, 4, 6, 8, 10}

	// WaitGroup позволяет дождаться выполнения всех горутин
	var wg sync.WaitGroup
	for _, v := range l {
		wg.Add(1)
		go printSquare(v, &wg)
	}
	wg.Wait()
	fmt.Println()
}

func printSquare(v int, wg *sync.WaitGroup) {
	fmt.Print(v*v, " ")
	wg.Done()
}
