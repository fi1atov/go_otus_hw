package main

import (
	"fmt"
	"sync"
)

func counterInGorutines(iteraionsNumber int) int {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	counter := 0

	for i := 0; i < iteraionsNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
		}()
	}
	wg.Wait()
	return counter
}

func main() {
	res := counterInGorutines(1000)
	fmt.Println(res)
}
