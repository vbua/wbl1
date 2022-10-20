package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	c := make(chan int64)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ctx context.Context, c chan int64, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down writer")
				return
			case c <- time.Now().Unix():
			}
			time.Sleep(1 * time.Second)
		}
	}(ctx, c, &wg)

	go func(ctx context.Context, c chan int64, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down reader")
				return
			case data := <-c:
				fmt.Println(data)
			}
		}
	}(ctx, c, &wg)

	wg.Wait()
	fmt.Println("Finished all goroutines")
}
