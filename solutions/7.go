package solutions

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Map struct {
	mRW  map[string]int
	lock sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		mRW:  make(map[string]int, 100),
		lock: sync.RWMutex{},
	}
}

// для безопасного чтения и записи используем RWMutex
func (m *Map) Read(key string) (int, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	res, ok := m.mRW[key]
	return res, ok
}

func (m *Map) Write(key string, value int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.mRW[key] = value
}

func Example7() {
	m := NewMap()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go workerMapRead(ctx, m)
	go workerMapWrite(ctx, m)

	select {
	case <-ctx.Done():
		fmt.Println("Finish")
	}
}

func workerMapRead(ctx context.Context, m *Map) {
	for {
		str := "abcdef"
		min, max := 0, len(str)-1
		c := rand.Intn(max-min) + min
		r := string(str[c])
		v, ok := m.Read(r)
		fmt.Printf("%s: %d - %t\n", r, v, ok)
	}
}

func workerMapWrite(ctx context.Context, m *Map) {
	for {
		str := "abcdef"
		min, max := 0, len(str)-1
		c := rand.Intn(max-min) + min
		m.Write(string(str[c]), rand.Int())
	}
}
