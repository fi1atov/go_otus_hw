package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writeInChannel(channel chan int) {
	for {
		n := rand.Intn(1000)
		// runtime.Gosched()
		time.Sleep(1 * time.Second)
		channel <- n
		fmt.Println("Insert to channel: ", n)
	}
}

func calculateAverage(list []int) int {
	total := 0
	for _, num := range list {
		total += num
	}
	return total / len(list)
}

func processingData(channelData, channelProcessed chan int) {
	for {
		dataSlice := make([]int, 0, 10)
		for s := range channelData {
			dataSlice = append(dataSlice, s)
			fmt.Println("Append 1 elem into slice: ", s)

			if len(dataSlice) == 10 {
				fmt.Println("10 elems reached")
				break
			}
		}
		fmt.Println("Get from channel 10 elems into slice: ", dataSlice)
		time.Sleep(1 * time.Second)
		average := calculateAverage(dataSlice)
		fmt.Println("Calculate average for 10 elems: ", average)
		channelProcessed <- average
		fmt.Println("Average into new channel: ", average)
	}
}

func printAverage(channelProcessed chan int) {
	for {
		fmt.Println("Printing calculate Average: ", <-channelProcessed)
	}
}

func main() {
	SensorDataChannel := make(chan int, 10)
	ProcessedDataChannel := make(chan int)

	go writeInChannel(SensorDataChannel)

	go processingData(SensorDataChannel, ProcessedDataChannel)

	go printAverage(ProcessedDataChannel)

	time.Sleep(60 * time.Second)
}
