package solutions

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Example4() {
	var n int
	fmt.Println("n: ")
	fmt.Fscan(os.Stdin, &n)
	ch := make(chan int, n)

	sigs := make(chan os.Signal, 1)
	// при получении CTRL+C записываем в канал
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// WaitGroup позволяет дождаться выполнения всех горутин
	var wg sync.WaitGroup
	wg.Add(n + 1)
	go workerPut(ch, &wg)
	for i := 0; i < n; i++ {
		go workerRead(i, ch, &wg)
	}

	// ожидаем запись в канал
	<-sigs
	fmt.Println("Stopping workers")
}

func workerRead(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Printf("Worker %d get: %d\n", i, <-ch)
	}
}

func workerPut(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ch <- rand.Int()
	}
}
