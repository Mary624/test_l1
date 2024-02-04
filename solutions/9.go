package solutions

import (
	"fmt"
	"sync"
)

func Example9() {
	chx := make(chan int)
	chx2 := make(chan int)
	l := []int{1, 2, 3, 4, 5}
	// WaitGroup позволяет дождаться выполнения всех горутин
	var wg sync.WaitGroup
	wg.Add(3)
	go getX(&wg, l, chx)
	go x2(&wg, chx, chx2)
	go x2Print(&wg, chx2)
	wg.Wait()
}

func getX(wg *sync.WaitGroup, l []int, chx chan int) {
	for _, v := range l {
		chx <- v
	}
	close(chx)
	wg.Done()
}

func x2(wg *sync.WaitGroup, chx, chx2 chan int) {
	for {
		x, ok := <-chx
		if ok {
			chx2 <- x * 2
			continue
		}
		close(chx2)
		wg.Done()
		return
	}
}

func x2Print(wg *sync.WaitGroup, chx2 chan int) {
	for {
		x2, ok := <-chx2
		if ok {
			fmt.Println(x2)
			continue
		}
		wg.Done()
		return
	}
}
