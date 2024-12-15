package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA bool

func Sensor(data chan<- float64) {
	cnt := 0
	for {
		if CLOSEA{
			return
		}
		data <- rand.Float64() * 100
		time.Sleep(100 * time.Millisecond)	
		cnt++	
		if cnt == 20{
			close(data)

			CLOSEA = true
		}
	}

}

func GetData(data <-chan float64, processedData chan<- float64) {
	var sum float64
	cnt := 0
	for value := range data{
		// if CLOSEA {
		// 	break
		// }
		sum += value
		cnt++
		if cnt == 10 {
			processedData <- sum / float64(cnt)
			sum = 0
			cnt = 0
		}
	}
	close(processedData)

	// for {
	// 	if CLOSEA {
	// 		close(processedData)
	// 		return
	// 	}
	// 	value := <-data
	// 	sum += value
	// 	cnt++
	// 	if cnt == 10 {
	// 		processedData <- sum / float64(cnt)
	// 		sum = 0
	// 		cnt = 0
	// 	}
	// }
}

func Info(processedData <-chan float64) {
	for mean := range processedData{
		fmt.Println(mean)
	}
	// for {
	// 	fmt.Println(<-processedData)
	// }
}

func main() {
	data := make(chan float64)
	processedData := make(chan float64)
	

	go Info(processedData)
	go GetData(data, processedData)
	go Sensor(data)
	select{
	
	}


	// time.Sleep(1 * time.Minute)
}
