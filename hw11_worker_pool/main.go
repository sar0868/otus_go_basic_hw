package main

import (
	"fmt"
	"sync"
)

var (
	cnt int
	mx  sync.Mutex
)

func Counter(n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		mx.Lock()
		cnt++
		mx.Unlock()
	}
	fmt.Printf("Counter %d completed work, cnt=%d done\n", n, cnt)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Counter(i, &wg, &mx)
	}
	wg.Wait()
}
