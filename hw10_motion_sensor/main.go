package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Sensor(data chan<- float64, wg *sync.WaitGroup) {
	timer := time.NewTimer(time.Minute)
label:
	for {
		select {
		case data <- rand.Float64() * 100:
			time.Sleep(500 * time.Millisecond)
		case <-timer.C:
			break label
		}
	}
	close(data)
	wg.Done()
}

func GetData(data <-chan float64, processedData chan<- float64, wg *sync.WaitGroup) {
	var sum float64
	cnt := 0
	for value := range data {
		sum += value
		cnt++
		if cnt == 10 {
			processedData <- sum / float64(cnt)
			sum = 0
			cnt = 0
		}
	}
	wg.Done()
	close(processedData)
}

func Info(processedData <-chan float64, wg *sync.WaitGroup) {
	for mean := range processedData {
		fmt.Println(mean)
	}
	wg.Done()
}

func main() {
	data := make(chan float64)
	processedData := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(3)
	go Info(processedData, &wg)
	go GetData(data, processedData, &wg)
	go Sensor(data, &wg)
	wg.Wait()
}
