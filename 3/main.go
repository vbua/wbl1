package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sumOfSquares(numbers []int) {
	wg := sync.WaitGroup{}
	var sum int32 = 0
	for _, v := range numbers {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			// используем атомик, потому что иначе будет рейс кондишен. ну и атомик будет быстрее мьютекса
			atomic.AddInt32(&sum, int32(v*v))
		}(v)
	}
	wg.Wait()
	fmt.Println(sum)
}

func main() {
	sumOfSquares([]int{2, 4, 6, 8, 10})
}
