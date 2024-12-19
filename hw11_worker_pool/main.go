package main

import (
	"fmt"
	"sync"
)

var (
	cnt int
	mx  sync.Mutex
)

func Counter(n int, wg *sync.WaitGroup, mx *sync.Mutex, ch chan string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		mx.Lock()
		cnt++
		mx.Unlock()
	}
	ch <- fmt.Sprintf("Counter %d completed work, cnt=%d done", n, cnt)
}

func Info(ch chan string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(<-ch)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Counter(i, &wg, &mx, ch)
		wg.Add(1)
		go Info(ch, &wg)
	}


	wg.Wait()
}
