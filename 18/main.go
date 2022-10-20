package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	number int64
}

func main() {
	counter := Counter{
		number: 0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			atomic.AddInt64(&c.number, 1)
		}(&counter, &wg)
	}
	wg.Wait()
	fmt.Println(counter.number)
}
