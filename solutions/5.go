package solutions

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Example5() {
	var n int
	fmt.Println("n: ")
	fmt.Fscan(os.Stdin, &n)
	ch := make(chan int, n)
	// контекст с таймаутом позволяет завершиться горутинам через n секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()
	fmt.Println("Start")
	go workerPut5(ctx, ch)
	go workerRead5(ctx, ch)
	select {
	case <-ctx.Done():
		fmt.Println("Finish")
	}
}

func workerRead5(ctx context.Context, ch chan int) {
	for {
		fmt.Printf("Worker get: %d\n", <-ch)
	}
}

func workerPut5(ctx context.Context, ch chan int) {
	for {
		ch <- rand.Int()
	}
}
