package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writeInChannel(channel chan int) {
	for {
		n := rand.Intn(1000)
		time.Sleep(1 * time.Second) // Достаточно только здесь добавить чтобы замедлить программу
		channel <- n
		// fmt.Println("Insert to channel: ", n)
	}
}

func calculateAverage(channel chan int) int {
	total := 0
	counter := 0
	for num := range channel {
		total += num
		counter++
		if counter == 10 {
			break
		}
	}
	return total / 10
}

func processingData(channelData, channelProcessed chan int) {
	for {
		if len(channelData) == 10 {
			// fmt.Println("10 elems received")
			average := calculateAverage(channelData)
			// fmt.Println("get average: ", average)
			channelProcessed <- average
		}
	}
}

func main() {
	SensorDataChannel := make(chan int, 10)
	ProcessedDataChannel := make(chan int)

	go writeInChannel(SensorDataChannel)

	go processingData(SensorDataChannel, ProcessedDataChannel)

	for {
		fmt.Println("Printing calculate Average: ", <-ProcessedDataChannel)
	}

	// Ждем 60 секунд, затем отправляем сигнал об остановке
	// Но сюда я никогда не попаду ведь у меня выше бесконечный цикл
	// time.Sleep(60 * time.Second)
	// stop <- true
}
