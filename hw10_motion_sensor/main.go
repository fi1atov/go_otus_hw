package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sensor() chan float64 {
	res := make(chan float64)
	timer := time.After(time.Minute)
	go func() {
		defer close(res)
		for {
			select {
			case res <- rand.Float64():
			case <-timer:
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return res
}

func SensorReader(sensorData <-chan float64) chan float64 {
	res := make(chan float64)

	go func() {
		defer close(res)
		i := 0
		var accum float64
		for v := range sensorData {
			i++
			accum += v
			if i%10 == 0 {
				res <- accum / 10.0
				accum = 0.0
			}
		}
	}()

	return res
}

func main() {
	dataCh := Sensor()
	acumCh := SensorReader(dataCh)

	for d := range acumCh {
		fmt.Println(d)
	}
}
