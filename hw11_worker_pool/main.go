package main

import (
	"fmt"
	"log"
	"sync"
)

func counterInGorutines(iteraionsNumber int) int {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	counter := 0

	for i := 0; i < iteraionsNumber; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
			log.Printf("Gorutine %d - task done\n", i)
		}(i)
	}
	wg.Wait()
	return counter
}

func main() {
	res := counterInGorutines(100)
	fmt.Println(res)
}
