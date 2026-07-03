package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	mu := sync.Mutex{}
	count := 0

	for range 100_000 {
		wg.Go(func() {
			mu.Lock()
			count2 := count
			count = count2 + 1
			mu.Unlock()
		})
	}

	wg.Wait()
	log.Println("count:", count)
}
