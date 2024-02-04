package solutions

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Counter struct {
	i    int
	lock sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		i:    0,
		lock: sync.Mutex{},
	}
}

// с помощью Mutex безопасно изменяем i
func (c *Counter) Add() {
	c.lock.Lock()
	c.i++
	c.lock.Unlock()
}

func (c *Counter) String() string {
	return strconv.Itoa(c.i)
}

func Example18() {
	c := NewCounter()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 4; i++ {
		go func(ctx context.Context) {
			for {
				c.Add()
			}
		}(ctx)
	}
	select {
	case <-ctx.Done():
		fmt.Println(c)
	}
}
