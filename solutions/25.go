package solutions

import (
	"context"
	"fmt"
	"time"
)

func Example25() {
	fmt.Println("Sleep...")
	sleep(1)
	fmt.Println("Wake up")
}

func sleep(sec int) {
	// timeout позволяет подождать n секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		for {

		}
	}(ctx)
	select {
	case <-ctx.Done():
	}
}
