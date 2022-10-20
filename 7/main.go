package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	storage := make(map[int]int)

	// создаем 11 горутин, конкурентно записывающи в мапу значения
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(storage map[int]int, mu *sync.Mutex, wg *sync.WaitGroup, i int) {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}(storage, &mu, &wg, i)
	}

	wg.Wait()
	fmt.Println(storage)
}
