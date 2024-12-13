package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sensor(data chan<- float64) {
	for {
		data <- rand.Float64() * 100
		time.Sleep(250 * time.Millisecond)
	}
}

func GetData(data <-chan float64, processedData chan<- float64) {
	var sum float64
	cnt := 0
	for {
		value := <-data
		sum += value
		cnt++
		if cnt == 10 {
			processedData <- sum / float64(cnt)
			sum = 0
			cnt = 0
		}
	}
}

func Info(processedData <-chan float64) {
	for {
		fmt.Println(<-processedData)
	}
}

func main() {
	data := make(chan float64)
	processedData := make(chan float64)

	go Info(processedData)
	go GetData(data, processedData)
	go Sensor(data)

	time.Sleep(1 * time.Minute)
}
