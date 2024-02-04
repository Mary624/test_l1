package solutions

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Example6() {
	// WaitGroup позволяет дождаться выполнения всех горутин
	var wg sync.WaitGroup
	wg.Add(1)
	// при выходе из функции горутина завершается
	go worker1(&wg)
	ch := make(chan bool)

	wg.Add(1)
	// используем канал и select для завершения
	go worker2(ch, &wg)
	time.Sleep(time.Second)
	ch <- true

	wg.Add(1)
	// паника завершает горутину
	go worker3(&wg)
	time.Sleep(time.Second)

	// контекст с cancel/timeout/deadline также завершает горутину
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker456(ctx, 4, &wg)
	time.Sleep(time.Second)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	wg.Add(1)
	defer cancel()
	go worker456(ctx, 5, &wg)
	time.Sleep(2 * time.Second)

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	wg.Add(1)
	defer cancel()
	go worker456(ctx, 6, &wg)
	time.Sleep(2 * time.Second)

	wg.Wait()
}

func worker1(wg *sync.WaitGroup) {
	fmt.Println("Finish worker 1")
	wg.Done()
}

func worker2(ch chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			fmt.Println("Finish worker 2")
			wg.Done()
			return
		}
	}
}

func worker3(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Finish worker 3")
			wg.Done()
		}
	}()
	panic("!")
}

func worker456(ctx context.Context, i int, wg *sync.WaitGroup) {
	defer func() {
		fmt.Printf("Finish worker %d\n", i)
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
